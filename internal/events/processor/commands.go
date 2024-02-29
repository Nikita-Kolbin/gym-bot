package processor

import (
	"log"
)

const (
	HelpCmd  = "/help"
	HelloCmd = "/hello"
)

func (p *Processor) doCmd(text string, chatID int, username string) error {
	log.Printf("got new command %s from %s", text, username)

	switch text {
	case HelpCmd:
		return p.sendHelp(chatID)
	case HelloCmd:
		return p.sendHello(chatID)
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
