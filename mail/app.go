package mail

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"
	"time"

	"github.com/BurntSushi/toml"
	_ "github.com/BurntSushi/toml"
)

type Config struct {
	SMTP SmtpConfig
}

var cfg *Config

func init() {

	log.Println("mail init")

	dir, err := os.Getwd()
	if err != nil {
		log.Println(err)
		panic(err)
	}

	configPath := path.Join(dir, "config", "server.toml")

	cfg = &Config{}
	if _, err := toml.DecodeFile(configPath, cfg); err != nil {
		log.Println(err)
		panic(err)
	}

}

func mailCheck() {

	log.Println(cfg)

	sender := MailSender{}
	sender.Config = cfg.SMTP

	message := MailMessage{}

	message.SetFrom(cfg.SMTP.EmailSenderName, cfg.SMTP.EmailSender)
	message.SetTo("", "")
	message.Subject = "Test Sub ject "
	message.IsHtml = true

	if fileHtml, err := os.ReadFile("files/F001.html"); err != nil {
		message.Body = err.Error()
	} else {
		// message.Body = string(fileHtml)

		messageTemplate := string(fileHtml)
		systemEmail := sender.Config.EmailSender
		occurDate := time.Now().Format("2006-01-02 15:04:05")

		target := make(map[string]interface{})

		target["os_ip_addr"] = ""
		target["threshold_type_name"] = "타입"
		target["code_name"] = "severity"
		messageContent := "empty "

		messageTemplate = strings.Replace(messageTemplate, "<?=email?>", systemEmail, 1)
		messageTemplate = strings.Replace(messageTemplate, "<?=occur_date?>", occurDate, 1)
		messageTemplate = strings.Replace(messageTemplate, "<?=ip_addr?>", target["os_ip_addr"].(string), 1)
		messageTemplate = strings.Replace(messageTemplate, "<?=category?>", target["threshold_type_name"].(string), 1)
		messageTemplate = strings.Replace(messageTemplate, "<?=severity?>", target["code_name"].(string), 1)
		messageTemplate = strings.Replace(messageTemplate, "<?=message?>", messageContent, 1)

		message.Body = messageTemplate
	}

	if err := sender.Send(&message); err != nil {
		fmt.Println(err)
	}

}

func Main() {

	mailCheck()

}
