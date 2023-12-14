package api

import (
	"bufio"
	"log"
	"os"
	"strings"
	"time"

	"github.com/spf13/viper"
)

func init() {
	log.Println("api test init.")
}

type SimpleSNS struct {
	Sender chan SimpleMessage
}

func (s *SimpleSNS) Run() {
	s.Sender = make(chan SimpleMessage)

	api := &Telegram{
		Token: viper.GetString("telegram.token"),
	}

	api.Init()

	go api.Listen()

	go func() {
		for m := range s.Sender {
			api.Send(m)
		}
	}()

	go func() {
		for m := range api.Receiver {
			log.Println(m)
		}
	}()

	s.readCommand()

}

func (s *SimpleSNS) readCommand() {
	reader := bufio.NewReader(os.Stdin)
	for {
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		log.Println(text)
		if strings.TrimSpace(text) == "quit" {
			os.Exit(1)
		}
		s.Sender <- SimpleMessage{
			Id:        1643426517,
			MessageId: int(time.Now().Unix()),
			Text:      text,
		}
		time.Sleep(1 * time.Second)
	}
}
