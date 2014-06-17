package main

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	// "github.com/mholt/binding"
	"fmt"
	"log"
	"net/http"
	// "github.com/unrolled/render"
	"github.com/wendal/goweixin"
	"github.com/younglucky/ohoh/models"
)

const (
	WX_TOKEN = "tk_wx_young"
)

var (
	router *mux.Router
)

func init() {
	log.Println("main.go init()...")
	goweixin.DevMode = true
	goweixin.SetDebug(true)

	router = mux.NewRouter()
}

func main() {
	router.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(rw, "hello world!")
	})
	router.Handle("/wx", &goweixin.WxHttpHandler{WX_TOKEN, &models.WeixinHandler{}})

	n := negroni.Classic()
	n.UseHandler(router)
	n.Run(":3000")
}
