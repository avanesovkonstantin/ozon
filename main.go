package main

import (
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

	for update := range updates {
		if update.Message != nil { // If we got a message

			switch update.Message.Command() {
			case "help":
				helpComand(bot, update.Message)
			case "list":
				listComand(bot, update.Message, productServise)
			default:
				defaultSend(bot, update.Message)
			}
		}
	}
}

func helpComand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "/help - help\n"+
		"/list - list of products")
	bot.Send(msg)
}

func listComand(bot *tgbotapi.BotAPI, message *tgbotapi.Message, productServise *product.Service) {
	allProducts := productServise.List()
	outputMessage := "Here are all the products:\n\n"
	for _, p := range allProducts {
		outputMessage += p.Title + "\n"
	}
	msg := tgbotapi.NewMessage(message.Chat.ID, outputMessage)
	bot.Send(msg)
}

func defaultSend(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Not recognise: "+message.Text)
	bot.Send(msg)
}
