package models

import (
	"github.com/wendal/goweixin"
	"log"
)

type WeixinHandler struct {
	*goweixin.BaseWeiXinHandler
}

func (w *WeixinHandler) Event(msg goweixin.Message) (reply goweixin.Replay) {
	log.Println("event:", msg.Event())
	return goweixin.ReplyText("欢迎关注，hello world!")
}
