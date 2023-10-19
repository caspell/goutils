package main

import (
	"fmt"
	_ "runtime/debug"
	_ "sync"
	"time"

	"log"
	// "github.com/wunicorns/goutils/batch"
	// "github.com/wunicorns/goutils/patterns"
	"github.com/wunicorns/goutils/channel"
)

func init() {
	fmt.Println("main package initialized. ", time.Now())
	log.Println("")
}

func main() {
	channel.Sample2()

}
