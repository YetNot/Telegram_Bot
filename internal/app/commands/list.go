package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (c *Commander) List(inputMessage *tgbotapi.Message) {
	outPutMsgText := "Here all the products: \n\n"

	products := c.productService.List()
	for _, p := range products {
		outPutMsgText += p.Title
		outPutMsgText += "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outPutMsgText)
	c.bot.Send(msg)
}
