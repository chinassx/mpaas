package task

import (
	"time"

	"github.com/infraboard/mpaas/apps/job"
)

func NewTaskSet() *JobTaskSet {
	return &JobTaskSet{
		Items: []*JobTask{},
	}
}

func (s *JobTaskSet) Add(task *JobTask) {
	s.Items = append(s.Items, task)
}

func NewDefaultTask() *JobTask {
	req := NewRunJobRequest("")
	return NewJobTask(req)
}

func NewJobTask(req *RunJobRequest) *JobTask {
	return &JobTask{
		Spec:   req,
		Job:    nil,
		Status: NewJobTaskStatus(),
	}
}

func (t *JobTask) Update(job *job.Job, status *JobTaskStatus) {
	t.Job = job
	t.Status = status
}

func NewJobTaskStatus() *JobTaskStatus {
	return &JobTaskStatus{
		StartAt: time.Now().Unix(),
	}
}
