package telegram

import "LinkStrashBot/clients/telegram"

type Processor struct {
	tg     *telegram.Client
	offset int
	//storage
}

func New()
