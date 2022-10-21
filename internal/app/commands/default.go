package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) DefaultSend(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Not recognise: "+message.Text)
	c.bot.Send(msg)
}
