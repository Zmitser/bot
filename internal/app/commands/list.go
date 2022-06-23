package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func (c *Commander) List(inputMessage *tgbotapi.Message) {
	outputMessage := "Here all products \n\n"
	for _, p := range c.service.List() {
		outputMessage += p.Title + "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMessage)
	c.bot.Send(msg)
}
