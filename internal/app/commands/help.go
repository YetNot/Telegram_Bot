package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (c *Commander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help - помощь\n"+
			"/list - список объектов\n"+
			"/get - получение объекта из списка\n"+
			"/new - создание нового объекта\n"+
			"/edit - изменение объекта\n"+
			"/delete - удалить объект",
	)
	c.bot.Send(msg)
}
