package utils

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

var CommandMenu = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("/random"),
		tgbotapi.NewKeyboardButton("/qst"),
		tgbotapi.NewKeyboardButton("/xp"),
	),
)

var ModeMenu = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Quiz"),
		tgbotapi.NewKeyboardButton("Decode"),
		tgbotapi.NewKeyboardButton("Memory"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Dare"),
		tgbotapi.NewKeyboardButton("Spy"),
	),
)

var ExitMenu = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("/exit"),
	),
)
