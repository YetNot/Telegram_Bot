package commands

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong args", args)
		return
	}

	product, err := c.productService.Describe(idx)
	if err != nil {
		log.Printf("all to get product %d: %v", idx, err)
		return
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("successfully parsed argument: %v", product.Title))
	//msg.ReplyToMessageID = update.Message.MessageID
	c.bot.Send(msg)
}