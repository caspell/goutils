package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

type Task struct {
	Job  []int
	Name string
}

func (t Task) Shout(idx int) int {
	fmt.Printf("%s: %d\n", t.Name, idx)
	return idx
}

type SampleRequest struct {
	Id      string
	Max     int
	Counter int
}

func (t *SampleRequest) exec(callback func(int)) {
	for t.Max > t.Counter {
		callback(t.Counter)
		t.Counter = t.Counter + 1
		time.Sleep(time.Second)
	}
}

func receive(q chan (int)) {

}

func sampleRun() {
	request := SampleRequest{"", 5, 0}

	q := make(chan int, 10000)

	go request.exec(func(num int) {
		fmt.Printf("in exec %d\n", num)
		q <- num
	})

	time.Sleep(time.Second * 3)
	fmt.Println(len(q))

	// var wg sync.WaitGroup
	// defer wg.Wait()

	// for {
	// 	values := <- q

	// 	wg.Add(1)
	// 	go func() {

	// 		fmt.Printf("out %d\n", 1)
	// 		// q <- num
	// 		fmt.Println(len(q))
	// 		fmt.Println("done")
	// 		fmt.Println("--------------------")

	// 		time.Sleep(time.Second * 3)

	// 		wg.Done()
	// 	}()

	// }

}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	// values := []string{"test1", "test2", "test3", "test4", "test5"}
	// async(values[0])
	sampleRun()

}

func async(v string) {

	workItems := make([]int, 30)

	var wg sync.WaitGroup
	defer wg.Wait()

	for item, _ := range workItems {
		// Wait until a worker is available
		wg.Add(1)
		go func(item int) {
			defer wg.Done()
			// time.Sleep(time.Second)
			fmt.Println(item)
		}(item)
	}

}

func async2(v string) {

	numWorkers := 10

	workItems := make([]int, 30)

	var wg sync.WaitGroup
	defer wg.Wait()

	pool := make(chan Task, numWorkers)

	for i := 0; i < numWorkers; i++ {
		pool <- Task{
			Name: fmt.Sprintf("Runner %s %d", v, i),
			Job:  make([]int, 10),
		}
	}

	for item, _ := range workItems {
		// Wait until a worker is available
		t := <-pool
		wg.Add(1)
		go func(item int) {
			defer wg.Done()
			time.Sleep(time.Second)
			t.Shout(item)
			pool <- t
		}(item)
		time.Sleep(1)
	}

}
