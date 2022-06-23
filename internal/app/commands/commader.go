package commands

import (
	"github.com/Zmitser/bot/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var registeredCommands = map[string]func(c *Commander, msg *tgbotapi.Message){}

type Commander struct {
	bot     *tgbotapi.BotAPI
	service *product.Service
}

func NewCommander(bot *tgbotapi.BotAPI, service *product.Service) *Commander {
	return &Commander{
		bot,
		service,
	}
}

func (c *Commander) HandleUpdate(update tgbotapi.Update) {
	if update.Message == nil { // ignore any non-Message Updates
		return
	}
	command, ok := registeredCommands[update.Message.Command()]

	if ok {
		command(c, update.Message)

	} else {
		c.Default(update.Message)
	}
}
