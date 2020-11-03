package main

import (
	"log"
	"net/http"

	"github.com/line/line-bot-sdk-go/linebot"
)

const channel_secret = "c8654b200557eb6744b21fc78f240d0a"
const channel_token = "G0k5a+PTDx8UU7s09xB3qD8viE7+AXGjrTJDBhiZ8Qs7j44nIKzlBgb0WFEqW7trLbJWcsX4HDXQJc3Bn+vJylb7S1sUnbzXiCWE9D4rUu9gsvYVVugCW4wc3dm6yQzlXsoPpBWIW4Kn1xZkOqRqjAdB04t89/1O/w1cDnyilFU="

func main() {
	bot, _ := linebot.New(channel_secret, channel_token)

	http.HandleFunc("/callback", func(w http.ResponseWriter, req *http.Request) {
		events, err := bot.ParseRequest(req)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				w.WriteHeader(400)
			} else {
				w.WriteHeader(500)
			}
			return
		}

		for _, event := range events {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				resp := linebot.NewTextMessage(message.Text)
				_, err := bot.ReplyMessage(event.ReplyToken, resp).Do()
				if err != nil {
					log.Print(err)
				}
			}
		}
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
