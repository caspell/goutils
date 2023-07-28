package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

func loop1(loopCount int) {

	now := time.Now()

	fmt.Println("loop count: ", loopCount)

	var wg sync.WaitGroup
	var m *sync.Map = &sync.Map{}

	defer wg.Wait()

	defer func() {
		deferTime := time.Since(now).Microseconds()
		fmt.Println(deferTime)
	}()

	for i := 0; i < loopCount; i++ {
		wg.Add(1)
		go func(j int) {
			m.Store(j, fmt.Sprintf("test %v", j))
			wg.Done()
		}(i)
	}

	// for i := 0; i < loopCount; i++ {
	// 	t, _ := m.Load(i)
	// 	fmt.Println("for loop: ", t)
	// }

	// m.Range(func(k, v interface{}) bool {
	// 	fmt.Println("range (): ", k, v)
	// 	return true
	// })

}

func loop2(loopCount int) {

	now := time.Now()

	fmt.Println("loop count: ", loopCount)

	var sm sync.Mutex

	// sm.TryLock()

	var wg sync.WaitGroup

	var m = make(map[int]string)

	defer wg.Wait()

	defer func() {
		deferTime := time.Since(now).Microseconds()
		fmt.Println(deferTime)
	}()

	for i := 0; i < loopCount; i++ {
		wg.Add(1)
		sm.Lock()
		go func(j int) {
			m[j] = fmt.Sprintf("test %v", j)
			sm.Unlock()
			wg.Done()
		}(i)
	}

	// for i := 0; i < loopCount; i++ {
	// 	// fmt.Println("for loop: ", t)
	// }

	sm.Lock()
	for i, v := range m {
		fmt.Println("range (): ", i, v)
	}
	sm.Unlock()

}

func main() {

	loopCount, _ := strconv.Atoi(os.Args[1])

	// fmt.Println(" first run ")
	// loop1(loopCount)

	fmt.Println(" second run ")
	loop2(loopCount)

}
