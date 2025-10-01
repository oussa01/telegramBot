package handlers

import (
	"math/rand"
	"time"
	"secSender/m/v2/models"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func SendChallengeForMode(bot *tgbotapi.BotAPI, user *models.User, mode string, chatID int64, exitKeyboard tgbotapi.ReplyKeyboardMarkup) {
	if !user.CanPlayToday() {
		bot.Send(tgbotapi.NewMessage(chatID, "⏳ You’ve reached your daily limit!"))
		return
	}

	if user.CurrentChallenge != nil {
		msg := tgbotapi.NewMessage(chatID, "❗ You are already in a challenge. Type /exit to quit it.")
		bot.Send(msg)
		return
	}

	challenges := models.FilterByType(models.Challenges, user.Answered, mode)
	if len(challenges) == 0 {
		bot.Send(tgbotapi.NewMessage(chatID, "✅ You’ve completed all "+mode+" challenges!"))
		return
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	challenge := challenges[r.Intn(len(challenges))]
	user.CurrentChallenge = &challenge
	user.Attempts = 0

	msg := tgbotapi.NewMessage(chatID, "🕹️ "+mode+" Challenge:\n"+challenge.Question)
	msg.ReplyMarkup = exitKeyboard
	bot.Send(msg)
}

func SendRandomChallenge(bot *tgbotapi.BotAPI, user *models.User, chatID int64, exitKeyboard tgbotapi.ReplyKeyboardMarkup) {
	if !user.CanPlayToday() {
		bot.Send(tgbotapi.NewMessage(chatID, "⏳ You’ve reached your daily limit!"))
		return
	}

	if user.CurrentChallenge != nil {
		msg := tgbotapi.NewMessage(chatID, "❗ You are already in a challenge. Type /exit to quit it.")
		bot.Send(msg)
		return
	}

	challenge := models.RandomChallenge(models.Challenges, user.Answered)
	user.CurrentChallenge = &challenge
	user.Attempts = 0

	msg := tgbotapi.NewMessage(chatID, "🎲 Random Challenge ("+challenge.Type+"):\n"+challenge.Question)
	msg.ReplyMarkup = exitKeyboard
	bot.Send(msg)
}
