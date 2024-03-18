package processor

import (
	"encoding/json"
	"gym-bot/internal/client"
)

const (
	kbdCreate         = "Создать"
	kbdCreateGroup    = "Создать группу"
	kbdCreateExercise = "Создать упражнение"

	kbdCancel = "Отмена"
)

func (p *Processor) sendDefaultKeyboard(chatID int, text string) error {
	k := client.ReplyKeyboardMarkup{Keyboard: make([][]client.KeyboardButton, 0)}
	k.Keyboard = append(k.Keyboard, []client.KeyboardButton{
		{Text: kbdCreate},
	})
	k.Keyboard = append(k.Keyboard, []client.KeyboardButton{
		{Text: "Button 2"},
	})
	k.Keyboard = append(k.Keyboard, []client.KeyboardButton{
		{Text: "Button 3"},
	})

	jKey, _ := json.Marshal(k)

	// TODO: do refactor

	return p.client.SendReplyKeyboardMarkup(chatID, string(jKey), text)
}

func (p *Processor) sendPickDefaultKeyboard(chatID int, text string) error {
	k := client.ReplyKeyboardMarkup{Keyboard: make([][]client.KeyboardButton, 0)}
	k.Keyboard = append(k.Keyboard, []client.KeyboardButton{
		{Text: kbdCreateGroup},
	})
	k.Keyboard = append(k.Keyboard, []client.KeyboardButton{
		{Text: kbdCreateExercise},
	})
	k.Keyboard = append(k.Keyboard, []client.KeyboardButton{
		{Text: kbdCancel},
	})

	jKey, _ := json.Marshal(k)

	// TODO: do refactor

	return p.client.SendReplyKeyboardMarkup(chatID, string(jKey), text)
}
