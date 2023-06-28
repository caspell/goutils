package batch

import (
	"fmt"
	"math/rand"
	"runtime"
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

func Module() {

	s := Scheduler{}
	s.Tasks = &[]Task{
		Task{
			Name:     "job1",
			Schedule: `*/6 * * * * *`,
			Handler:  job1(rand.Int()),
		},
		Task{
			Name:     "job2",
			Schedule: `*/3 * * * * *`,
			Handler:  job2(rand.Int()),
		},
	}

	s.Start()

	fmt.Println("done!")
	time.Sleep(time.Second * 3)
}
