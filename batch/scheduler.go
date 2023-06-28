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
		var entryId cron.EntryID
		var err error
		if t.Handler != nil {
			entryId, err = c.AddFunc(t.Schedule, t.Handler)
		} else if t.Job != nil {
			entryId, err = c.AddFunc(t.Schedule, t.Job)
		}
		if err != nil {
			t.Error = err
		} else {
			t.EntryId = entryId
		}
	}

	c.Start()

}
