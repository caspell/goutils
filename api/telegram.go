package api

import (
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Telegram struct {
	Token  string
	ChatId int64
	*tgbotapi.BotAPI
	Receiver chan SimpleMessage
}

type SimpleMessage struct {
	Id        int64
	MessageId int
	Name      string
	Text      string
	DateTime  time.Time
}

func (t *Telegram) Init() {
	bot, err := tgbotapi.NewBotAPI(t.Token)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	t.BotAPI = bot
	t.Receiver = make(chan SimpleMessage)
	t.ChatId = t.lookupChat()
}

func (t *Telegram) GetChatId() (int64, error) {

	var chatId int64
	var err error

	u := tgbotapi.NewUpdate(0)

	u.Timeout = 60

	updates, err := t.BotAPI.GetUpdates(u)
	if err != nil {
		log.Println(err)
	}

	log.Println(updates)

	if len(updates) > 0 {
		chatId = updates[len(updates)-1].Message.Chat.ID
	} else {
		log.Println("no messages")
	}

	return chatId, err
}

func (t *Telegram) Send(msg SimpleMessage) error {
	botMsg := tgbotapi.NewMessage(t.ChatId, msg.Text)
	rst, err := t.BotAPI.Send(botMsg)
	if err != nil {
		return err
	}
	log.Println(rst)
	return nil
}

func (t *Telegram) lookupChat() int64 {
	log.Printf("Authorized on account %s", t.BotAPI.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := t.BotAPI.GetUpdatesChan(u)
	for update := range updates {
		if update.Message != nil {
			return update.Message.Chat.ID
		}
	}
	return -1
}

func (t *Telegram) Listen() {

	log.Printf("Authorized on account %s", t.BotAPI.Self.UserName)

	u := tgbotapi.NewUpdate(0)

	u.Timeout = 60

	updates := t.BotAPI.GetUpdatesChan(u)

	for update := range updates {

		if update.Message != nil { // If we got a message

			log.Printf("AAA: [%s] %s", update.Message.From.UserName, update.Message.Text)
			log.Println("AAA:: ", update.Message.Chat)

			t.Receiver <- SimpleMessage{
				Id:        update.Message.Chat.ID,
				MessageId: update.Message.MessageID,
				Name:      update.Message.From.UserName,
				Text:      update.Message.Text,
				DateTime:  time.Unix(int64(update.Message.Date), 0),
			}
		}
	}
}
