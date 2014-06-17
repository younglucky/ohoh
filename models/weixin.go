package models

import (
	"encoding/json"
	"github.com/wendal/goweixin"
	"io/ioutil"
	"log"
	"os"
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
		SaveUser(msg)
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

func SaveUser(msg goweixin.Message) {
	bytes, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	fileName := "./data/user.json"
	err = ioutil.WriteFile(filename, bytes, os.ModePerm)
	if err != nil {
		panic(err)
	}

}

func main() {
	msg := goweixin.Message{"hello": "world"}
	SaveUser(msg)
}
