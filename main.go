package main

import (
	"fmt"
	_ "runtime/debug"
	_ "sync"
	"time"

	"log"
	// "github.com/wunicorns/goutils/batch"
	// "github.com/wunicorns/goutils/patterns"
)

func init() {
	fmt.Println("main package initialized. ", time.Now())
	log.Println("")
}

type MainRef struct {
	Name   string
	Ints   []int
	Floats *[]float64
}

type MainObject struct {
	Name string
	Ref1 *MainRef
	Ref2 MainRef
}

func main() {

}
