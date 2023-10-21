package commands

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) List(inputMessage *tgbotapi.Message) {
	outPutMsgText := "Here all the products: \n\n"

	products := c.productService.List()
	for _, p := range products {
		outPutMsgText += fmt.Sprintf("%v: ", p.ID)
		outPutMsgText += p.Title
		outPutMsgText += "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outPutMsgText)

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("next page", "some data"),
		),
	)

	c.bot.Send(msg)
}
