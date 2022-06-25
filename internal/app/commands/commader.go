package commands

import (
	"github.com/Zmitser/bot/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
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
	defer func() {
		if panicValue := recover(); panicValue != nil {
			log.Printf("Recover from panic value: %v", panicValue)
		}
	}()

	if update.CallbackQuery != nil {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Data: "+update.CallbackQuery.Message.Text)
		c.bot.Send(msg)
		return
	}

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
