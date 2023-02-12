package impl

import (
	"context"
	"fmt"
	"time"

	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mpaas/apps/pipeline"
	"github.com/infraboard/mpaas/apps/task"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// 执行Pipeline
func (i *impl) RunPipeline(ctx context.Context, in *task.RunPipelineRequest) (
	*task.PipelineTask, error) {
	// 1. 查询需要执行的Pipeline
	p, err := i.pipeline.DescribePipeline(ctx, pipeline.NewDescribePipelineRequest(in.Id))
	if err != nil {
		return nil, err
	}
	ins := task.NewPipelineTask(p)

	// 从pipeline 取出需要执行的任务
	t := ins.GetFirstJobTask()
	if t == nil {
		return nil, fmt.Errorf("not job task to run")
	}

	// 保存Job Task, 所有JobTask 批量生成, 全部处于Pendding状态, 然后入库 等待状态更新
	err = i.JobTaskBatchSave(ctx, ins.JobTasks())
	if err != nil {
		return nil, err
	}

	// 运行 第一个Job, 驱动Pipeline执行
	ins.MarkedRunning()
	resp, err := i.RunJob(ctx, t.Spec)
	if err != nil {
		return nil, err
	}
	t.Update(resp.Job, resp.Status)

	// 保存Pipeline状态
	if _, err := i.pcol.InsertOne(ctx, ins); err != nil {
		return nil, exception.NewInternalServerError("inserted a pipeline task document error, %s", err)
	}

	return ins, nil
}

// Pipeline中任务有变化时,
// 如果执行成功则 继续执行, 如果失败则标记Pipeline结束
// 当所有任务成功结束时标记Pipeline执行成功
func (i *impl) PipelineTaskStatusChanged(ctx context.Context, in *task.JobTask) (
	*task.PipelineTask, error) {
	if in == nil && in.Status == nil {
		return nil, exception.NewBadRequest("job task or job task status is nil")
	}

	if in.Spec.PipelineTask == "" {
		return nil, exception.NewBadRequest("Pipeline Id参数缺失")
	}

	// 获取Pipeline Task, 因为Job Task是先保存在触发的回调, 这里获取的Pipeline Task是最新的
	descReq := task.NewDescribePipelineTaskRequest(in.Spec.PipelineTask)
	p, err := i.DescribePipelineTask(ctx, descReq)
	if err != nil {
		return nil, err
	}

	switch in.Status.Stage {
	case task.STAGE_PENDDING, task.STAGE_ACTIVE, task.STAGE_CANCELING:
		// Pipeline Task状态无变化
		return p, nil
	case task.STAGE_CANCELED:
		// 任务取消, pipeline 取消执行
		p.MarkedCanceled()
		return i.updatePipelineStatus(ctx, p)
	case task.STAGE_FAILED:
		// 任务执行失败, 更新Pipeline状态为失败
		if !in.Spec.IgnoreFailed {
			p.MarkedFailed()
			return i.updatePipelineStatus(ctx, p)
		}
	case task.STAGE_SUCCEEDED:
		// 任务运行成功, pipeline继续执行
	}

	// task执行成功或者忽略执行失败, pipeline 仍然处于运行中, 需要获取下一个任务执行
	nexts := p.NextRun()

	// 如果没有需要执行的任务, Pipeline执行结束, 更新Pipeline状态为成功
	if nexts == nil && nexts.Len() == 0 {
		p.MarkedSuccess()
		return i.updatePipelineStatus(ctx, p)
	}

	// 如果有需要执行的JobTask, 继续执行
	for index := range nexts.Items {
		item := nexts.Items[index]
		t, err := i.RunJob(ctx, item.Spec)
		if err != nil {
			return nil, err
		}
		item.Status = t.Status
		item.Job = t.Job
	}

	return p, nil
}

// 更新Pipeline状态
func (i *impl) updatePipelineStatus(ctx context.Context, in *task.PipelineTask) (*task.PipelineTask, error) {
	in.Status.UpdateAt = time.Now().Unix()
	if _, err := i.pcol.UpdateByID(ctx, in.Meta.Id, bson.M{"$set": bson.M{"status": in.Status}}); err != nil {
		return nil, exception.NewInternalServerError("update task(%s) document error, %s",
			in.Meta.Id, err)
	}
	return in, nil
}

// 查询Pipeline任务
func (i *impl) QueryPipelineTask(ctx context.Context, in *task.QueryPipelineTaskRequest) (
	*task.PipelineTaskSet, error) {
	r := newQueryPipelineTaskRequest(in)
	resp, err := i.pcol.Find(ctx, r.FindFilter(), r.FindOptions())

	if err != nil {
		return nil, exception.NewInternalServerError("find pipeline task error, error is %s", err)
	}

	set := task.NewPipelineTaskSet()
	// 循环
	for resp.Next(ctx) {
		ins := task.NewDefaultPipelineTask()
		if err := resp.Decode(ins); err != nil {
			return nil, exception.NewInternalServerError("decode pipeline task  error, error is %s", err)
		}
		set.Add(ins)
	}

	// count
	count, err := i.pcol.CountDocuments(ctx, r.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("get pipeline task count error, error is %s", err)
	}
	set.Total = count
	return set, nil
}

// 查询Pipeline任务详情
func (i *impl) DescribePipelineTask(ctx context.Context, in *task.DescribePipelineTaskRequest) (
	*task.PipelineTask, error) {
	filter := bson.M{"_id": in.Id}

	ins := task.NewDefaultPipelineTask()
	if err := i.pcol.FindOne(ctx, filter).Decode(ins); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("pipeline task %s not found", in.Id)
		}

		return nil, exception.NewInternalServerError("find pipeline task %s error, %s", in.Id, err)
	}

	// 补充该PipelineTask管理的JobTask
	query := task.NewQueryTaskRequest()
	query.PipelineTaskId = in.Id
	tasks, err := i.QueryJobTask(ctx, query)
	if err != nil {
		return nil, err
	}

	// 将tasks 填充给pipeline task
	for i := range tasks.Items {
		t := tasks.Items[i]
		stage := ins.GetStage(t.Spec.StageName)
		if stage != nil {
			stage.Add(t)
		}
	}

	return ins, nil
}

// 删除Pipeline任务详情
func (i *impl) DeletePipelineTask(ctx context.Context, in *task.DeletePipelineTaskRequest) (
	*task.PipelineTask, error) {
	ins, err := i.DescribePipelineTask(ctx, task.NewDescribePipelineTaskRequest(in.Id))
	if err != nil {
		return nil, err
	}
	// TODO: 运行中的流水线不运行删除, 先取消 才能删除

	// 删除该Pipeline下所有的Job Task
	tasks := ins.Status.JobTasks()
	for index := range tasks.Items {
		t := tasks.Items[index]

		// 没有运行过的任务不需要清理
		if t.Status.Stage.Equal(task.STAGE_PENDDING) {
			continue
		}

		_, err := i.DeleteJobTask(ctx, task.NewDeleteJobTaskRequest(t.Meta.Id))
		if err != nil {
			if !exception.IsNotFoundError(err) {
				return nil, err
			}
		}
	}

	if err := i.deletecluster(ctx, ins); err != nil {
		return nil, err
	}

	return ins, nil
}
