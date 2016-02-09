package main

import "time"
import "github.com/tucnak/telebot"
import "os"

func main() {
	telegramApiKey := os.Getenv("DJIGURBOT_TELEGRAMAPIKEY")
	if telegramApiKey == "" {
		panic("Telegram api key is not set.")
	}

	bot, err := telebot.NewBot(telegramApiKey)
	if err != nil {
		panic("Could not connect to Telegram API")
	}

	messages := make(chan telebot.Message)
	bot.Listen(messages, 1*time.Second)

	for message := range messages {
		if message.Text == "/hi" {
			bot.SendMessage(message.Chat,
				"Привет, "+message.Sender.FirstName+"!", nil)
		}
	}
}
