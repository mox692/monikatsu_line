package server

import (
	"fmt"
	"log"
	"net/http"
	"regexp"

	"monikatsuline/constant"

	"github.com/line/line-bot-sdk-go/linebot"
	"golang.org/x/xerrors"
)

const channel_secret = "c8654b200557eb6744b21fc78f240d0a"
const channel_token = "G0k5a+PTDx8UU7s09xB3qD8viE7+AXGjrTJDBhiZ8Qs7j44nIKzlBgb0WFEqW7trLbJWcsX4HDXQJc3Bn+vJylb7S1sUnbzXiCWE9D4rUu9gsvYVVugCW4wc3dm6yQzlXsoPpBWIW4Kn1xZkOqRqjAdB04t89/1O/w1cDnyilFU="

type AppErr struct {
	status string
	err    error
}

// bot, _ := linebot.New(channel_secret, channel_token)
type LineConn struct {
	bot    *linebot.Client
	event  *linebot.Event
	events []*linebot.Event
}

func Serve() {
	http.HandleFunc("/callback", JudgeEvent)
}

// JudgeEventはLINE MessagingAPIから渡されたイベントを判別します。
// 後続処理はそれぞれのイベントの関数に移譲させます。
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

// judgeContextは、テキストメッセージを送ってきたuserが会話の途中かどうかを判別します。
func (c *LineConn) judgeContext(message *linebot.TextMessage) {

	// kvsマイクロサービスにアクセスし、セッションを確認
	status := getContext(c.event.Source.UserID)

	// status毎に、その後の処理を切り分け
	switch status[:0] {
	// default
	case "0":
		c.defaultContact(message)
	// 登録
	case "1":

	// モニカツ登録
	case "2":
		switch status[1:2] {
		case "1":
			c.setWakeupTime(message)
		case "2":
		}

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
	err := setContext(c.event.Source.UserID, "2.1")
	if err != nil {
		log.Print(err)
	}

	resp := linebot.NewTextMessage(constant.MONIKATSU_RESISTER)
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

/* monikatu関連 */

// setWakeupTime
func (c *LineConn) setWakeupTime(message *linebot.TextMessage) {

	// 入力メッセージをvalidate
	time, err := parseInputTime(message.Text)
	if err != nil {
		resp := linebot.NewTextMessage(constant.MONIKATSU_WRONGTIME_INPUT)
		_, err = c.bot.ReplyMessage(c.event.ReplyToken, resp).Do()
		if err != nil {
			log.Print(err)
		}
		return
	}
	fmt.Println(time)

	// kvsマイクロサービスにセッションをinsert
	err = setContext(c.event.Source.UserID, "2.2")
	if err != nil {
		log.Print(err)
	}

	resp := linebot.NewTextMessage(constant.MONIKATSU_WAKEUP_RESISTER(time))
	_, err = c.bot.ReplyMessage(c.event.ReplyToken, resp).Do()
	if err != nil {
		log.Print(err)
	}
}

// デモ関数。実際はここはgRPCのメソッドに挿しかわる。
func getContext(userid string) string {
	return ""
}

// デモ関数。実際はここはgRPCのメソッドに挿しかわる。
func setContext(userid, contextID string) error {
	return nil
}

func parseInputTime(time string) (string, error) {
	r, _ := regexp.Compile("^([0-9]|1[0-9]|2[0-3]|):[0-5][0-9]")
	if !(r.Match([]byte(time))) {
		return "", xerrors.New("parse err!")
	}
	return time, nil
}
