package main

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func main2() {
	bot, err := tgbotapi.NewBotAPI("")
	if err != nil {
		log.Panic(err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		chatID := update.Message.Chat.ID

		reply := tgbotapi.NewMessage(chatID, fmt.Sprintf("Received message from : %d", chatID))
		_, err := bot.Send(reply)
		if err != nil {
			log.Println(err)
		}
	}
}
