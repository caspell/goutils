package main

import (
	"fmt"
	"sync"
)

func Run() {

	wg := &sync.WaitGroup{}

	mu := &sync.Mutex{}

	for i := 0; i < 20; i++ {
		wg.Add(1)
		mu.Lock()
		go func(index int) {
			defer func() {
				wg.Done()
				mu.Unlock()
			}()
			fmt.Println(index)
		}(i)
	}

	wg.Wait()

	fmt.Println("done!")
}
