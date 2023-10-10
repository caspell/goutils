package batch

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

type Accumulator struct {
}

func (a Accumulator) NewJob() func() {

	return func() {

	}
}

func Goroutines() {
	b := make([]byte, 1024)
	size := runtime.Stack(b, false)
	fmt.Println(size, string(b))
}

func job1(num int) func() {
	fmt.Println("job1 ", num)

	return func() {
		fmt.Printf("in job1 %d, %d \n", num, rand.Int())
		Goroutines()
		fmt.Println("----------------------------------- 1")
	}
}

func job2(num int) func() {
	fmt.Println("job2 ", num)

	return func() {
		fmt.Printf("in job2 %d ,%d \n", num, rand.Int())
		Goroutines()
		fmt.Println("----------------------------------- 2")
	}
}

func TestRunner() {

	now := time.Now()

	fmt.Println("Entry!", now.UnixMilli())

	wg := &sync.WaitGroup{}

	for i := 0; i < 7; i++ {
		wg.Add(1)
		go func(index int) {
			fmt.Println("Print!", index)
			time.Sleep(1 * time.Second)
			wg.Done()
		}(i)
	}

	wg.Wait()

	fmt.Println("Exit!")

}

func TestRunnerBackup() {

	err := recover()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("done!")
}

func TestRunner2() {
	defer TestRunnerBackup()
	value := 0
	if value < 1 {
		panic("value zero!!!!!")
	}
	fmt.Println(100 / value)
}

func Module() {

	s := Scheduler{}
	s.Tasks = &[]Task{
		// Task{
		// 	Name:     "job1",
		// 	Schedule: `*/6 * * * * *`,
		// 	Handler:  job1(rand.Int()),
		// },
		// Task{
		// 	Name:     "job2",
		// 	Schedule: `*/3 * * * * *`,
		// 	Handler:  job2(rand.Int()),
		// },
		// Task{
		// 	Name:     "job3",
		// 	Schedule: `*/10 * * * * *`,
		// 	Handler:  TestRunner,
		// },
		Task{
			Name:     "job4",
			Schedule: `*/5 * * * * *`,
			Handler:  TestRunner2,
		},
	}

	s.Start()

	fmt.Println("DONE!")

}
