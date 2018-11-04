package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/line/line-bot-sdk-go/linebot"
)

const envAccessToken = "LB_ACCESS_TOKEN"
const envSecretToken = "LB_SECRET_TOKEN"
const envPort = "LB_PORT"
const envLiffURL = "LB_LIFF_URL"

var bot *linebot.Client
var liffURL string

func init() {
	st := os.Getenv(envSecretToken)
	at := os.Getenv(envAccessToken)
	botTmp, err := linebot.New(st, at)
	if err != nil {
		panic("init error")
	}

	bot = botTmp
	liffURL = os.Getenv(envLiffURL)
}

func sorryMessage(event *linebot.Event, w http.ResponseWriter) error {
	message := linebot.NewTextMessage("Sorry!! Not Supported.")
	_, err := bot.ReplyMessage(event.ReplyToken, message).Do()
	if err != nil {
		log.Println(err)
	}
	return nil
}

func liffLink(event *linebot.Event, w http.ResponseWriter) error {
	uriAction := linebot.NewURIAction("リンク", liffURL)
	bt := linebot.NewButtonsTemplate("", "LIFF", "以下よりxxxをしてください", uriAction)
	message := linebot.NewTemplateMessage("LIFF Button Template", bt)
	_, err := bot.ReplyMessage(event.ReplyToken, message).Do()
	if err != nil {
		log.Println(err)
	}
	return nil
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	events, err := bot.ParseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for _, event := range events {
		switch event.Type {
		case linebot.EventTypeMessage:
			tm, ok := event.Message.(*linebot.TextMessage)
			if !ok {
				log.Println("not text message")
				continue
			}

			if strings.Contains(tm.Text, "ページ") {
				liffLink(event, w)
				continue
			}

			sorryMessage(event, w)
		case linebot.EventTypeBeacon:
		default:
			log.Printf("not supported type: %v", event.Type)
		}
	}
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func run() error {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/ping", pingHandler)

	port := os.Getenv(envPort)
	fmt.Printf("START: port=%v\n", port)
	return http.ListenAndServe(":"+port, nil)
}

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}
