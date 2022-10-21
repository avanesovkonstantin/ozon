package commands

import (
	"ak/ozon/tgbot/internal/service/product"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Commander struct {
	bot            *tgbotapi.BotAPI
	productServise *product.Service
}

func NewCommander(bot *tgbotapi.BotAPI, productServise *product.Service) *Commander {
	return &Commander{
		bot:            bot,
		productServise: productServise,
	}
}
