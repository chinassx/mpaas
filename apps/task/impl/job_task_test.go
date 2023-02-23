package impl_test

import (
	"os"
	"testing"

	"github.com/infraboard/mpaas/apps/job"
	"github.com/infraboard/mpaas/apps/pipeline"
	"github.com/infraboard/mpaas/apps/task"
	"github.com/infraboard/mpaas/test/conf"
	"github.com/infraboard/mpaas/test/tools"
)

func TestRunBuildJob(t *testing.T) {
	req := pipeline.NewRunJobRequest("docker_build@default.default")
	version := job.NewVersionedRunParam("v1")
	version.Params = job.NewRunParamWithKVPaire(
		"GIT_ADDRESS", "git@github.com:infraboard/mpaas.git",
		"IMAGE_REPOSITORY", "registry.cn-hangzhou.aliyuncs.com/inforboard/mpaas",
		"IMAGE_VERSION", "v0.0.1",
	)
	req.Params = version

	ins, err := impl.RunJob(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(ins))
}

func TestRunDeployJob(t *testing.T) {
	req := pipeline.NewRunJobRequest("docker_deploy@default.default")
	version := job.NewVersionedRunParam("v1")
	version.Params = job.NewRunParamWithKVPaire(
		job.SYSTEM_VARIABLE_DEPLOY_CONFIG_ID, conf.C.DEPLOY_CONFIG_ID,
	)
	req.Params = version

	ins, err := impl.RunJob(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(ins))
}

func TestQueryJobTask(t *testing.T) {
	req := task.NewQueryTaskRequest()
	req.PipelineTaskId = os.Getenv("PIPELINE_TASK_ID")
	set, err := impl.QueryJobTask(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(tools.MustToJson(set))
}

func TestDescribeJobTask(t *testing.T) {
	req := task.NewDescribeJobTaskRequest(os.Getenv("JOB_TASK_ID"))
	ins, err := impl.DescribeJobTask(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(ins))
}

func TestUpdateJobTaskStatus(t *testing.T) {
	req := task.NewUpdateJobTaskStatusRequest(os.Getenv("JOB_TASK_ID"))
	req.Stage = task.STAGE_SUCCEEDED
	req.Message = "执行成功"
	req.Detail = tools.MustReadContentFile("test/k8s_job.yml")
	ins, err := impl.UpdateJobTaskStatus(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToYaml(ins))
}

func TestDeleteJobTask(t *testing.T) {
	req := task.NewDeleteJobTaskRequest(os.Getenv("JOB_TASK_ID"))
	set, err := impl.DeleteJobTask(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(set)
}
