package translator

import (
	"github.com/fujiwaram/gopherTranslator/consts"
	"github.com/fujiwaram/gopherTranslator/translator/lib"
)

type message struct {
	msg string
}

func newMessage(msg string) message {
	return message{msg: msg}
}

func (msg message) translate() *lib.Order {
	var o lib.Order
	o.Printf("%s %s\n", consts.OrderSay, msg.msg)
	return &o
}
