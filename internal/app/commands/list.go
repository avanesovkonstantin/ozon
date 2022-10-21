package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) List(message *tgbotapi.Message) {
	allProducts := c.productServise.List()
	outputMessage := "Here are all the products:\n\n"
	for _, p := range allProducts {
		outputMessage += p.Title + "\n"
	}
	msg := tgbotapi.NewMessage(message.Chat.ID, outputMessage)
	c.bot.Send(msg)
}
