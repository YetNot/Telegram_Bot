package commands

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) New(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()
	newObject, err := c.productService.Create(args)
	if err != nil {
		log.Printf("all to get product %d: %v", newObject, err)
	}
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Succesfull!")
	c.bot.Send(msg)
}