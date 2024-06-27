package process

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

func init() {
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(fmt.Sprintf("%s start", hostname))
}

func main() {

	args := os.Args

	log.Println(args)

	if sub() {
		subProcess()
	} else {
		Run(args)
	}

}

func sub() bool {
	var sub *bool = flag.Bool("sub", false, "")
	flag.Parse()
	return *sub
}

func Run(args []string) {
	log.Println("@ root process")
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)
	defer cancel()
	cmd := exec.CommandContext(ctx, args[0], "--sub")
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	log.Println("@ root process done")
}

func subProcess() {
	log.Println("@ sub process")
	time.Sleep(5 * time.Second)
	data2 := []byte(`test ok !!`)
	os.WriteFile("D:/test.log", data2, os.ModeAppend)
	log.Println("@ sub process done!")
}
