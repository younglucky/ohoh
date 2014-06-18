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
	e := msg.Event()
	ek := msg.EventKey()
	switch {
	case e == "unsubscribe": //取消订阅
		log.Println("取消关注")
		return w.Default(msg)
	case e == "subscribe":
		s := ""
		if ek == "" { //关注
			s = "查找关注"
		} else if strings.HasPrefix(ek, "qrscene_") { //未关注时的扫码关注
			s = "扫码关注"
		} else { //已关注后的扫码
			s = "你已关注"
		}
		SaveUser(msg)
		return goweixin.ReplyText("欢迎关注，hello world!" + s)
	}

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
