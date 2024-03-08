package processor

import (
	"encoding/json"
	"gym-bot/internal/client"
	"log"
	"strings"
)

const (
	StartCmd = "/start"
	HelpCmd  = "/help"
	TestCmd  = "/test"
)

func (p *Processor) doCmd(text string, chatID int, username string) error {
	text = strings.ToLower(text)

	log.Printf("got new command %s from %s", text, username)

	switch text {
	case HelpCmd:
		return p.sendHelp(chatID)
	case StartCmd:
		if err := p.storage.CreateUser(username); err != nil {
			return err
		}
		return p.sendHello(chatID)
	case TestCmd:
		return p.sendKeyboard(chatID)
	default:
		return p.sendUnknownCommand(chatID)
	}
}

func (p *Processor) sendHelp(chatID int) error {
	return p.client.SendMessage(chatID, msgHelp)
}

func (p *Processor) sendHello(chatID int) error {
	return p.client.SendMessage(chatID, msgHello)
}

func (p *Processor) sendUnknownCommand(chatID int) error {
	return p.client.SendMessage(chatID, msgUnknownCommands)
}

func (p *Processor) sendKeyboard(chatID int) error {
	k := client.ReplyKeyboardMarkup{Keyboard: make([][]client.KeyboardButton, 0)}
	k.Keyboard = append(k.Keyboard, []client.KeyboardButton{
		{Text: "Button 1 skidi wop wop wop yes yes"},
	})
	k.Keyboard = append(k.Keyboard, []client.KeyboardButton{
		{Text: "Button 2"},
	})
	k.Keyboard = append(k.Keyboard, []client.KeyboardButton{
		{Text: "Button 3"},
	})

	jkey, _ := json.Marshal(k)

	return p.client.SendReplyKeyboardMarkup(chatID, string(jkey))
}
