package processor

import (
	"fmt"
	"gym-bot/internal/storage"
)

func (p *Processor) handleMsg(text string, chatID int, username string) error {
	state, err := p.storage.CheckState(username)
	if err != nil {
		return nil
	}

	switch state {
	case storage.Standard:
		return p.handleStandardState(text, chatID, username)
	case storage.CreateGroup:
		return p.handleCrateGroupState(text, chatID, username)
	default:
		return p.sendUnknownMessage(chatID)
	}

	// TODO: log all actions
}

func (p *Processor) handleStandardState(text string, chatID int, username string) error {
	switch text {
	case kbdCreateGroup:
		err := p.storage.ChangeState(username, storage.CreateGroup)
		if err != nil {
			return err
		}
		return p.sendCreateGroup(chatID)
	default:
		return p.sendUnknownMessage(chatID)
	}
}

func (p *Processor) handleCrateGroupState(text string, chatID int, username string) error {
	if err := p.storage.CreateGroup(username, text); err != nil {
		return err
	}

	if err := p.storage.ChangeState(username, storage.Standard); err != nil {
		return err
	}

	if err := p.sendCreateGroupSuccess(chatID, text); err != nil {
		return err
	}

	return nil
}

func (p *Processor) sendCreateGroup(chatID int) error {
	return p.client.SendMessage(chatID, msgCreateGroup)
}

func (p *Processor) sendUnknownMessage(chatID int) error {
	return p.client.SendMessage(chatID, msgUnknownMessage)
}

func (p *Processor) sendCreateGroupSuccess(chatID int, text string) error {
	m := fmt.Sprintf(`Группа "%s" создана`, text)
	return p.client.SendMessage(chatID, m)
}
