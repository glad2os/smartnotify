package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
)

func Send(telegramId int64, messageText string) {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("API"))

	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	msg := tgbotapi.NewMessage(telegramId, messageText)

	if _, err := bot.Send(msg); err != nil {
		log.Panic(err)
	}
}
