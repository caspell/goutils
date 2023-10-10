package batch

import (
	cron "github.com/robfig/cron/v3"
)

type Task struct {
	Name     string
	Schedule string
	Handler  func()
	Job      cron.FuncJob
	EntryId  cron.EntryID
	Error    error
}

type Scheduler struct {
	Tasks *[]Task
}

func (s Scheduler) Start() {
	c := cron.New(cron.WithSeconds())
	for _, t := range *s.Tasks {
		if t.Handler != nil {
			t.EntryId, t.Error = c.AddFunc(t.Schedule, t.Handler)
		} else if t.Job != nil {
			t.EntryId, t.Error = c.AddFunc(t.Schedule, t.Job)
		}
	}
	c.Start()
}
