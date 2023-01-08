package adapters

import (
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelgramClient struct {
	Client *tgbotapi.BotAPI
	ChatID int64
}

func NewTelegramClient() (*TelgramClient, error) {
	client, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_API"))
	if err != nil {
		return nil, err
	}
	return &TelgramClient{Client: client, ChatID: -1001894162684}, nil
}

func (c *TelgramClient) SendToGroup(msg string) error {
	m := tgbotapi.NewMessage(c.ChatID, msg)
	_, err := c.Client.Send(m)
	return err
}
