package patterns

import (
	"log"
	"sync"
	"time"
)

type TaskTracker struct {
	TaskName        string
	StartTime       time.Time
	DoneTimes       []float64
	Callbacks       []func(self *TaskTracker)
	WaitGroup       *sync.WaitGroup
	TaskFailedCount uint
	TaskChannel     chan int
	ChannelClosed   bool
}

func MakeTracker(name string, isAsync bool, requiredCalls ...func(self *TaskTracker)) TaskTracker {
	tt := TaskTracker{
		TaskName:        name,
		StartTime:       time.Now(),
		TaskFailedCount: 0,
	}
	tt.WaitGroup = &sync.WaitGroup{}
	tt.DoneTimes = make([]float64, 1)
	tt.TaskChannel = make(chan int)
	tt.Callbacks = requiredCalls
	go tt.Watch()
	return tt
}

func (tt *TaskTracker) Wait() {
	tt.WaitGroup.Wait()
	close(tt.TaskChannel)
	tt.ChannelClosed = true
	tt.TaskChannel = nil
}

func (tt *TaskTracker) Add() {
	tt.WaitGroup.Add(1)
}

func (tt *TaskTracker) Done() {
	tt.DoneTimes = append(tt.DoneTimes, float64(time.Since(tt.StartTime).Milliseconds()))
	tt.WaitGroup.Done()
}

func (tt *TaskTracker) Watch() {
	for v := range tt.TaskChannel {
		tt.TaskFailedCount = tt.TaskFailedCount + uint(v)
	}
	tt.ChannelClosed = true
}

func (tt *TaskTracker) Notice(after ...func(self *TaskTracker)) {
	if err := recover(); err != nil {
		log.Println(tt.TaskName, err)
		if !tt.ChannelClosed {
			tt.TaskChannel <- 1
		}
	}
	for _, callback := range append(tt.Callbacks, after...) {
		defer callback(tt)
	}
	tt.Wait()
}

func (tt *TaskTracker) ForEach(datas []interface{}, executions ...func(interface{})) {
	for _, v := range datas {
		tt.Add()
		go func(args interface{}) {
			defer func() {
				tt.Done()
				if err := recover(); err != nil {
					log.Println(tt.TaskName, err)
					if !tt.ChannelClosed {
						tt.TaskChannel <- 1
					}
				}
			}()
			for _, execute := range executions {
				execute(args)
			}
		}(v)
	}
}
