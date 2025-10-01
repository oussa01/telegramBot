package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"secSender/m/v2/config"
	"secSender/m/v2/handlers"
	"secSender/m/v2/models"
	"secSender/m/v2/utils"
	"strconv"
	"strings"

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
	webhookURL := "telegrambot-production-4840.up.railway.app/" + bot.Token
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

// Get or create a user session
func getUser(userID int) *models.User {
	user, exists := users[userID]
	if !exists {
		user = &models.User{
			ID:       userID,
			Answered: make(map[string]bool),
			XP:       0,
			Attempts: 0,
		}
		users[userID] = user
	}
	return user
}

// Simple string matching
func matchAnswer(given, expected string) bool {
	return strings.TrimSpace(strings.ToLower(given)) == strings.TrimSpace(strings.ToLower(expected))
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
		handlers.SendChallengeForMode(bot, user, text, chatID, utils.ExitMenu)
		return
	}

	switch text {
	case "/qst":
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
		handlers.SendRandomChallenge(bot, user, chatID, utils.ExitMenu)
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
			msg := tgbotapi.NewMessage(chatID, "No active challenge. Use /qst or /random to start one.")
			msg.ReplyMarkup = utils.CommandMenu
			bot.Send(msg)
		}
		return
	}

	// Answer handling
	if user.CurrentChallenge != nil {
		if matchAnswer(text, user.CurrentChallenge.Answer) {
			user.Answered[user.CurrentChallenge.Question] = true
			user.XP += 10
			user.CurrentChallenge = nil
			user.Attempts = 0
			user.MarkPlayed()

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
			user.XP -= 1
			hintMsg := ""
			if user.Attempts >= 3 && user.CurrentChallenge.Hint != "" {
				hintMsg = "\n💡 Hint: " + user.CurrentChallenge.Hint + " (-2 XP)"
			}
			bot.Send(tgbotapi.NewMessage(chatID, "❌ Incorrect! Try again. (-1 XP)"+hintMsg))
		}
		return
	}

	// Default response
	msg := tgbotapi.NewMessage(chatID, "👉 Use /qst or /random to play! Max 3 missions/day.")
	msg.ReplyMarkup = utils.CommandMenu
	bot.Send(msg)
}
