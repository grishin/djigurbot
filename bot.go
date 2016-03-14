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
		} else if strings.Contains(strings.ToLower(message.Text), "костя") {
			bot.SendMessage(message.Chat, "Костя крутой!", nil)
		} else if strings.Contains(strings.ToLower(message.Text), "твоя!") {
			bot.SendMessage(message.Chat, "Нееееет, твоя!", nil)
		} else if strings.Contains(strings.ToLower(message.Text), "крутой") {
			bot.SendMessage(message.Chat, "Нееееет, "+message.Sender.FirstName+", это ты крутой!", nil)
		} else if strings.Contains(strings.ToLower(message.Text), "сокиабле") {
			bot.SendMessage(message.Chat, "Сокиабле? ЧОБЛЯ?", nil)
		} else if strings.Contains(strings.ToLower(message.Text), "доброе утро") {
			bot.SendMessage(message.Chat, "И тебе наидобрейшего утра, "+message.Sender.FirstName+"!", nil)
		} else if strings.Contains(strings.ToLower(message.Text), "поздравляй!") {
			bot.SendMessage(message.Chat, `С 8 марта поздравляем вас, коллеги,
От души хотим вам пожелать,
Чтоб совместные победы и успехи
Дали нам возможность процветать!
Чтоб в делах житейских и в работе
Находить умели компромисс,
Одевались по последней моде,
Были леди, то бишь миссис или мисс.
Чтобы было нам в кого влюбляться,
Чтобы было нас кому любить,
Молодыми вечно оставаться
и насыщенной веселой жизнью жить!`, nil)
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
