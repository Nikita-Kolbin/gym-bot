package processor

import (
	"encoding/json"
	"gym-bot/internal/client"
)

const (
	kbdCreateGroup = "Создать группу"
)

func (p *Processor) sendDefaultKeyboard(chatID int) error {
	k := client.ReplyKeyboardMarkup{Keyboard: make([][]client.KeyboardButton, 0)}
	k.Keyboard = append(k.Keyboard, []client.KeyboardButton{
		{Text: kbdCreateGroup},
	})
	k.Keyboard = append(k.Keyboard, []client.KeyboardButton{
		{Text: "Button 2"},
	})
	k.Keyboard = append(k.Keyboard, []client.KeyboardButton{
		{Text: "Button 3"},
	})

	jkey, _ := json.Marshal(k)

	// TODO: do refactor

	return p.client.SendReplyKeyboardMarkup(chatID, string(jkey))
}
