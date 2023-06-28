package main

import (
	"fmt"
	_ "sync"
	"time"
)

type TaskTime struct {
	Name string
	float64
}

func init() {
	fmt.Println("init")
}

func main() {

	ch := make(chan TaskTime)

	num := 10

	// var mutex = &sync.Mutex{}
	// mutex.Lock()
	// defer mutex.Unlock()

	for i := 0; i < num; i++ {
		go func(i int) {
			ch <- TaskTime{
				Name:    "value",
				float64: 10.0 * float64(i),
			}
		}(i)
	}

	go func() {
		fmt.Println("out")
		for {
			chv := <-ch
			fmt.Printf("%v \n", chv.float64)
		}
	}()

	time.Sleep(time.Second * 1)
}
