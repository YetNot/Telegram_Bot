package commands

import (
	"github.com/YetNot/bot/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Delete(inputMessage *tgbotapi.Message) {
	message := "Successfull!"
	args := inputMessage.CommandArguments()
	arg, err := product.Translate(args)
	if err != nil {
		message = "Надо другой формат"
	}
	found := c.productService.Remove(uint64(arg))
	if !found {
		message = "Такого id нет"
	}
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, message)
	c.bot.Send(msg)
}
