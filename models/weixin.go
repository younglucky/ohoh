package models

import (
	"encoding/json"
	"github.com/wendal/goweixin"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type WeixinHandler struct {
	*goweixin.BaseWeiXinHandler
}

func (w *WeixinHandler) Text(msg goweixin.Message) goweixin.Replay {
	c := strings.ToLower(msg.Content())
	log.Println("c:", c)
	switch c {
	case "goelia":
		return goweixin.ReplyText("http://ohoh.co/goelia/subscribe/news/" + msg.ToUserName())
	default:
		return goweixin.ReplyText("你说的\"" + msg.Content() + "\"" + "我们没找到相关内容")
	}

}
func (w *WeixinHandler) Image(msg goweixin.Message) goweixin.Replay {
	return goweixin.ReplyText("你发的是图片")
}
func (w *WeixinHandler) Location(msg goweixin.Message) goweixin.Replay {
	return w.Default(msg)
}
func (w *WeixinHandler) Link(msg goweixin.Message) goweixin.Replay {
	return w.Default(msg)
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
	return w.Default(msg)
}
func (w *WeixinHandler) Default(msg goweixin.Message) goweixin.Replay {
	return nil
}

func SaveUser(msg goweixin.Message) {
	bytes, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	filename := "./data/user.json"
	err = ioutil.WriteFile(filename, bytes, os.ModePerm)
	if err != nil {
		panic(err)
	}

}

func main() {
	msg := goweixin.Message{"hello": "world"}
	SaveUser(msg)
}
