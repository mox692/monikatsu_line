package server

import (
	"log"
	"net/http"

	"monikatsuline/constant"
	"monikatsuline/sessionClient"

	"github.com/line/line-bot-sdk-go/linebot"
	"golang.org/x/xerrors"
)

const channel_secret = "c8654b200557eb6744b21fc78f240d0a"
const channel_token = "G0k5a+PTDx8UU7s09xB3qD8viE7+AXGjrTJDBhiZ8Qs7j44nIKzlBgb0WFEqW7trLbJWcsX4HDXQJc3Bn+vJylb7S1sUnbzXiCWE9D4rUu9gsvYVVugCW4wc3dm6yQzlXsoPpBWIW4Kn1xZkOqRqjAdB04t89/1O/w1cDnyilFU="

type AppErr struct {
	status string
	err    error
}

// LineConn はMessagingAPIとの接続を管理する構造体です。
// MessagingAPIのイベントオブジェクトとClient構造体をラップしています。
type LineConn struct {
	bot    *linebot.Client
	event  *linebot.Event
	events []*linebot.Event
}

// SessionCode はsession codeを管理します。
type SessionCode string

var (
	DefaultState           SessionCode = "0"
	MonikatsuFlag          SessionCode = "2.1"
	MonikatsuSetWakeupTime SessionCode = "2.2"
)

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
func (c *LineConn) judgeContext(message *linebot.TextMessage) error {

	// kvsマイクロサービスにアクセスし、セッションを確認
	status, err := sessionClient.GetContext(c.event.Source.UserID)
	if err != nil {
		return xerrors.Errorf("sessionClient.GetContext err : %w", err)
	}

	// status毎に、その後の処理を切り分け
	switch status.Data[:0] {
	// default
	default:
		c.defaultContact(message)
	// 登録
	case "1":
		// registerインスタンスを作成
		r := new(resister)
		switch status.Data[1:2] {
		case "1":
			r.askAppName(message)
		case "2":
			r.askPassword(message)
		}
	// モニカツ登録
	case "2":
		m := new(monikatsu)
		switch status.Data[1:2] {
		case "1":
			m.setWakeupTime(message)
		case "2":
		}
	}
	return nil
}

// context statusが0の際の処理。
// フック単語がuserから発された場合、各botイベントのinit処理が呼ばれます。
func (c *LineConn) defaultContact(message *linebot.TextMessage) error {

	// コメントの解析からの,
	switch message.Text {
	case "モニカツ", "monikatsu", "もにかつ":
		err := c.initMonikatsu()
		if err != nil {
			return err
		}
	case "とうろく":
		err := c.initResister()
		if err != nil {
			return err
		}
	default:
		// ヘルプを表示。
		c.helpMessage()
	}
	// リプライの分岐
	return nil
}

// セッションにuseridを登録して、モニカツ予約フラグを立てる。
func (c *LineConn) initMonikatsu() error {
	// kvsマイクロサービスにセッションをinsert
	err := setContext(c.event.Source.UserID, "2.1")
	if err != nil {
		return err
	}

	resp := linebot.NewTextMessage(constant.MONIKATSU_RESISTER)
	_, err = c.bot.ReplyMessage(c.event.ReplyToken, resp).Do()
	if err != nil {
		return err
	}
	return nil
}

func (c *LineConn) helpMessage() {
	resp := linebot.NewTextMessage(constant.HelpMessage)
	_, err := c.bot.ReplyMessage(c.event.ReplyToken, resp).Do()
	if err != nil {
		log.Print(err)
	}
}

func (c *LineConn) initResister() error {
	err := setContext(c.event.Source.UserID, "1.1")
	if err != nil {
		return err
	}
	resp := linebot.NewTextMessage(constant.RESISTER_INIT)
	_, err = c.bot.ReplyMessage(c.event.ReplyToken, resp).Do()
	if err != nil {
		log.Print(err)
	}
	return nil
}

// デモ関数。実際はここはgRPCのメソッドに挿しかわる。
func GetContext(userid string) string {

	return ""
}

// デモ関数。実際はここはgRPCのメソッドに挿しかわる。
func setContext(userid, contextID string) error {
	return nil
}
