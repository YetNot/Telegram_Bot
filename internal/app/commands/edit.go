package commands

import (
	"log"
	"strings"

	"github.com/YetNot/bot/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Edit(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()
	arg := strings.Split(args, " ")
	id, err := product.Translate(arg[0])
	if err != nil {
		log.Println("Введи число! ", err)
		return
	}
	err = c.productService.Update(uint64(id), arg[1])
	if err != nil {
		log.Println("Ошибка!")
		return
	}
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Successfull!")
	c.bot.Send(msg)
}
