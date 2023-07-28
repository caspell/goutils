package patterns

import (
	"fmt"
	"sync"
)

func Wrapper(id string) func(string) {

	var mux = &sync.Mutex{}

	mux.Lock()

	fmt.Printf("locked : %s \n", id)

	return func(id2 string) {

		defer mux.Unlock()

		fmt.Printf("unlocked : %s %s \n", id, id2)

	}
}

func Execute(id string) {

	defer Wrapper(id)(id)

	fmt.Println("done")

}
