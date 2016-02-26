package main

import (
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/tucnak/telebot"
)

type toast struct {
	Text string
}

func main() {
	telegramAPIKey := os.Getenv("DJIGURBOT_TELEGRAMAPIKEY")
	if telegramAPIKey == "" {
		panic("Telegram api key is not set.")
	}

	bot, err := telebot.NewBot(telegramAPIKey)
	if err != nil {
		panic("Could not connect to Telegram API")
	}

	toasts, err := readToastsFile()
	if err != nil {
		panic("Could not read toasts.txt file")
	}

	messages := make(chan telebot.Message)
	bot.Listen(messages, 1*time.Second)

	for message := range messages {
		if message.Text == "/hi" {
			bot.SendMessage(message.Chat,
				"Привет, "+message.Sender.FirstName+"!", nil)
		} else if strings.Contains(message.Text, "Тост!") {
			randomToast := toasts[rand.Intn(len(toasts))]
			bot.SendMessage(message.Chat, randomToast.Text, nil)
			/*} else if message.Text == "Костя!" {
			bot.SendMessage(message.Chat, "Костя занят, он на речном!", nil) */
		} else if strings.Contains(message.Text, "Твоя!") {
			bot.SendMessage(message.Chat, "Нееееет, твоя!", nil)
		} else if strings.Contains(message.Text, "крутой") {
			bot.SendMessage(message.Chat, "Нееееет, "+message.Sender.FirstName+", это ты крутой!", nil)
		}
	}
}

func readToastsFile() ([]toast, error) {
	content, err := ioutil.ReadFile("toasts.txt")
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(content), "---")

	toasts := make([]toast, len(lines))
	for i, line := range lines {
		t := toast{Text: line}
		toasts[i] = t
	}

	return toasts, err
}
