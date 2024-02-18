package process

import (
	"log"
	"os"
	"os/signal"
)

func Daemonize() error {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill)
	for {
		select {
		case killSignal := <-interrupt:
			if killSignal == os.Interrupt {
				log.Println("os interrupt")
				return nil
			}
		}
	}
}

func Daemon(pidFile string) error {
	Pid := New(pidFile)
	if err := Daemonize(); err != nil {
		return err
	}
	if err := Pid.Free(); err != nil {
		return err
	}
	return nil
}
