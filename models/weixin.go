package models

import (
	"github.com/wendal/goweixin"
	"log"
)

type WeixinHandler struct {
	*goweixin.BaseWeiXinHandler
}

func (w *WeixinHandler) Text(msg goweixin.Message) goweixin.Replay {
	return h.Default(msg)
}
func (w *WeixinHandler) Image(msg goweixin.Message) goweixin.Replay {
	return h.Default(msg)
}
func (w *WeixinHandler) Location(msg goweixin.Message) goweixin.Replay {
	return h.Default(msg)
}
func (w *WeixinHandler) Link(msg goweixin.Message) goweixin.Replay {
	return h.Default(msg)
}
func (w *WeixinHandler) Event(msg goweixin.Message) goweixin.Replay {
	log.Println("event:", msg)
	if msg.Event() == "subscribe" {
		return goweixin.ReplyText("欢迎关注，hello world!")
	}
	return w.Default(msg)
}
func (w *WeixinHandler) Voice(msg goweixin.Message) goweixin.Replay {
	return h.Default(msg)
}
func (w *WeixinHandler) Default(msg goweixin.Message) goweixin.Replay {
	return nil
}
