package main

import (
	"ak/ozon/tgbot/internal/app/commands"
	"ak/ozon/tgbot/internal/service/product"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	token := os.Getenv("TOKEN")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	productServise := product.NewService()

	commander := commands.NewCommander(bot, productServise)

	for update := range updates {
		if update.Message != nil {

			switch update.Message.Command() {
			case "help":
				commander.Help(update.Message)
			case "list":
				commander.List(update.Message)
			default:
				commander.DefaultSend(update.Message)
			}
		}
	}
}
