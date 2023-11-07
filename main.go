package main

import (
	"embed"
	"fmt"
	_ "runtime/debug"
	_ "sync"
	"time"

	"github.com/wunicorns/goutils/mail"
	// "github.com/wunicorns/goutils/batch"
	// "github.com/wunicorns/goutils/channel"
	// "github.com/wunicorns/goutils/api"
)

//go:embed files
var d embed.FS

func init() {
	fmt.Println("main package initialized. ", time.Now())
}

func main() {

	// api.Run()
	// fi.Include(&d)

	mail.Main()

}
