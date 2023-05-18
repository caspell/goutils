package batch

import (
	cron "github.com/robfig/cron/v3"
)

type Task struct {
	Schedule string
	Handler  func()
	EntryId  cron.EntryID
	Error    error
}

type Scheduler struct {
	Tasks *[]Task
}

func (s Scheduler) Start() {

	c := cron.New(cron.WithSeconds())

	for _, t := range *s.Tasks {

		if entryId, err := c.AddFunc(t.Schedule, t.Handler); err != nil {
			t.Error = err
		} else {
			t.EntryId = entryId
		}

	}

	c.Start()

}
