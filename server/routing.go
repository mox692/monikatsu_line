package server

import (
	"log"
	"net/http"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/mox692/monikatsu_line/constant"
)

const channel_secret = "c8654b200557eb6744b21fc78f240d0a"
const channel_token = "G0k5a+PTDx8UU7s09xB3qD8viE7+AXGjrTJDBhiZ8Qs7j44nIKzlBgb0WFEqW7trLbJWcsX4HDXQJc3Bn+vJylb7S1sUnbzXiCWE9D4rUu9gsvYVVugCW4wc3dm6yQzlXsoPpBWIW4Kn1xZkOqRqjAdB04t89/1O/w1cDnyilFU="

// bot, _ := linebot.New(channel_secret, channel_token)
type LineConn struct {
	bot    *linebot.Client
	event  *linebot.Event
	events []*linebot.Event
}

func Serve() {
	http.HandleFunc("/callback", JudgeEvent)
}

func JudgeEvent(w http.ResponseWriter, r *http.Request) {

	bot, _ := linebot.New(channel_secret, channel_token)

	events, err := bot.ParseRequest(r)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}

	c := &LineConn{
		bot:    bot,
		events: events,
	}

	for _, event := range c.events {

		c.event = event

		// 現時点ではMessageイベントのみの処理を実装。
		switch message := event.Message.(type) {

		// Textmessageの場合
		case *linebot.TextMessage:
			c.judgeContext(message)
		}
	}
}

// テキストメッセージを送ってきたuserが、会話の途中かどうかを判別
func (c *LineConn) judgeContext(message *linebot.TextMessage) {

	// kvsマイクロサービスにアクセスし、セッションを確認
	status := getContext(c.event.Source.UserID)

	// status毎に、その後の処理を切り分け
	switch status {
	case "":
		// 大業種に振り分け。
	case "1":
		// デフォルト操作
	}

}

// context statusが0の際の処理
func (c *LineConn) defaultContact(message *linebot.TextMessage) {

	// コメントの解析からの,
	switch message.Text {
	case "モニカツ", "monikatsu", "もにかつ":
		c.resisterMonikatsu(message)
	default:
		// ヘルプを表示。
		c.helpMessage()
	}

	// リプライの分岐

}

// セッションにuseridを登録して、モニカツ予約フラグを立てる。
func (c *LineConn) resisterMonikatsu(message *linebot.TextMessage) {
	// kvsマイクロサービスにセッションをinsert
	err := setContext(c.event.Source.UserID)

	resp := linebot.NewTextMessage(message.Text)
	_, err = c.bot.ReplyMessage(c.event.ReplyToken, resp).Do()
	if err != nil {
		log.Print(err)
	}
}

func (c *LineConn) helpMessage() {
	resp := linebot.NewTextMessage(constant.HelpMessage)
	_, err := c.bot.ReplyMessage(c.event.ReplyToken, resp).Do()
	if err != nil {
		log.Print(err)
	}
}

// デモ関数。実際はここはgRPCのメソッドに挿しかわる。
func getContext(userid string) string {
	return ""
}

// デモ関数。実際はここはgRPCのメソッドに挿しかわる。
func setContext(userid string) error {
	return nil
}
