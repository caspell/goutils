package patterns

import (
	"log"
	"sync"
	"time"
	// "github.com/wunicorns/goutils/batch"
)

var taskTimeMap *sync.Map
var TaskChan chan Task
var RootChannel chan int

const (
	FailedMetricsSuffix = "Failed"
	MaxLoop             = 10
)

type Task struct {
	Name      string
	TimeSpent float64
	Failed    float64
}

func init() {
	log.Println("patterns package initialized.")
	taskTimeMap = &sync.Map{}
	TaskChan = make(chan Task)
	RootChannel = make(chan int)
	go Tracking()
}

// JobTrace(batch name) in Batch -> receive chan TimeTakeChan
func Tracking() {
	defer close(TaskChan)
	for {
		timeTracking := <-TaskChan
		taskName := timeTracking.Name
		ms := timeTracking.TimeSpent
		failed := timeTracking.Failed
		taskTimeMap.Store(taskName, ms)
		taskTimeMap.Store(taskName+FailedMetricsSuffix, failed)
		log.Println(taskName, ", time spent: ", ms, ", failed: ", failed)
	}
}

func CreateTaskTracker(name string) *TaskTracker {
	tt := &TaskTracker{
		TaskName:        name,
		StartTime:       time.Now(),
		WaitGroup:       &sync.WaitGroup{},
		TaskFailedCount: 0,
		TaskChannel:     make(chan int),
	}
	tt.Callbacks = []func(self *TaskTracker){
		func(self *TaskTracker) {
			TaskChan <- Task{
				Name:      self.TaskName,
				TimeSpent: float64(time.Since(tt.StartTime).Milliseconds()),
				Failed:    float64(tt.TaskFailedCount),
			}
		},
	}
	return tt
}

func void(num int) {

}

func Run2() {
	ZeroValue := 0

	tracker := CreateTaskTracker("test tracker 2")

	defer tracker.Notice()

	datas := make([]interface{}, 0)

	for i := 0; i < MaxLoop; i++ {
		data := make(map[string]int)

		data["index"] = i
		data["testkey2"] = i * 99

		datas = append(datas, data)
	}

	for _, v := range datas {
		d := v.(map[string]int)
		void(d["index"] * ZeroValue)
		time.Sleep(1 * time.Second)
	}

}

func Run1() {
	ZeroValue := 0

	tracker := CreateTaskTracker("test tracker 1")

	defer tracker.Notice()

	datas := make([]interface{}, 0)

	for i := 0; i < MaxLoop; i++ {
		data := make(map[string]int)

		data["index"] = i
		data["testkey2"] = i * 99

		datas = append(datas, data)
	}

	tracker.ForEach(datas, func(v interface{}) {

		d := v.(map[string]int)
		void(d["index"] * ZeroValue)
		time.Sleep(1 * time.Second)

	})
}
