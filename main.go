package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"secSender/m/v2/config"
	"secSender/m/v2/handlers"
	"secSender/m/v2/models"
	"secSender/m/v2/utils"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var users = make(map[int]*models.User)

func main() {
	cfg := config.LoadConfig()
	bot, err := tgbotapi.NewBotAPI(cfg.BotToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false

	// Webhook URL (Railway provides HTTPS)
	webhookURL := "https://YOUR_RAILWAY_APP_URL/" + bot.Token
	_, err = bot.SetWebhook(tgbotapi.NewWebhook(webhookURL))
	if err != nil {
		log.Panic(err)
	}

	http.HandleFunc("/"+bot.Token, func(w http.ResponseWriter, r *http.Request) {
		var update tgbotapi.Update
		if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
			log.Println("Error decoding update:", err)
			return
		}
		handleUpdate(bot, update)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Bot running on port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func handleUpdate(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	if update.Message == nil {
		return
	}

	userID := update.Message.From.ID
	chatID := update.Message.Chat.ID
	text := update.Message.Text

	if users[userID] == nil {
		users[userID] = &models.User{ID: userID, Answered: make(map[string]bool)}
	}
	user := users[userID]

	// Mode selection
	if user.SelectingMode {
		user.SelectingMode = false
		handlers.SendChallengeForMode(bot, user, text, chatID,utils.ExitMenu)
		return
	}

	switch text {
	case "/question":
		if user.CurrentChallenge != nil {
			bot.Send(tgbotapi.NewMessage(chatID, "❗ Finish your current challenge or type /exit first."))
			return
		}
		msg := tgbotapi.NewMessage(chatID, "Choose a mode:")
		msg.ReplyMarkup = utils.ModeMenu
		bot.Send(msg)
		user.SelectingMode = true
		return
	case "/random":
		handlers.SendRandomChallenge(bot, user, chatID,utils.ExitMenu)
		return
	case "/xp":
		bot.Send(tgbotapi.NewMessage(chatID, "⭐ Your XP: "+strconv.Itoa(user.XP)))
		return
	case "/exit":
		if user.CurrentChallenge != nil {
			user.CurrentChallenge = nil
			user.Attempts = 0
			msg := tgbotapi.NewMessage(chatID, "You exited the challenge. Back to the main menu.")
			msg.ReplyMarkup = utils.CommandMenu
			bot.Send(msg)
		} else {
			msg := tgbotapi.NewMessage(chatID, "No active challenge. Use /question or /random to start one.")
			msg.ReplyMarkup = utils.CommandMenu
			bot.Send(msg)
		}
		return
	}

	// Answer handling
	if user.CurrentChallenge != nil {
		if models.MatchAnswer(text, user.CurrentChallenge.Answer) {
			user.Answered[user.CurrentChallenge.Question] = true
			user.XP += 10
			user.MarkPlayed()
			user.CurrentChallenge = nil
			user.Attempts = 0

			progress := len(user.Answered)
			total := len(models.Challenges)

			msg := tgbotapi.NewMessage(chatID,
				"✅ Correct! +10 XP\n"+
					"You’ve completed "+strconv.Itoa(progress)+"/"+strconv.Itoa(total)+" challenges.\n"+
					"⭐ Total XP: "+strconv.Itoa(user.XP)+"\n"+
					"📅 Today’s missions: "+strconv.Itoa(user.DailyCount)+"/3")
			msg.ReplyMarkup = utils.CommandMenu
			bot.Send(msg)
		} else {
			user.Attempts++
			user.XP -= 2
			if user.Attempts >= 3 && user.CurrentChallenge.Hint != "" {
				bot.Send(tgbotapi.NewMessage(chatID, "❌ Wrong again! Hint: "+user.CurrentChallenge.Hint+" (-2 XP)"))
			} else {
				bot.Send(tgbotapi.NewMessage(chatID, "❌ Incorrect! Try again. (-2 XP)"))
			}
		}
		return
	}

	// Default response
	msg := tgbotapi.NewMessage(chatID, "👉 Use /question or /random to play! Max 3 missions/day.")
	msg.ReplyMarkup = utils.CommandMenu
	bot.Send(msg)
}
