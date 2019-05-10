package translator

import (
	"github.com/fujiwaram/gopher-translator/consts"
	"github.com/fujiwaram/gopher-translator/translator/lib"
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
