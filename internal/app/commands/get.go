package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
)

func (c *Commander) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	idx, err := strconv.Atoi(args)

	if err != nil {
		log.Println("wrong args", args)
	}

	product, err := c.service.Get(idx)

	if err != nil {
		log.Printf("failed to get product wuth id %d: %v", idx, args)
	}
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, product.Title)
	c.bot.Send(msg)
}

func init() {
	registeredCommands["get"] = (*Commander).Get
}
