package service

// import (
// 	"log"

// 	"monikatsuline/server"

// 	"monikatsuline/constant"

// 	"github.com/line/line-bot-sdk-go/linebot"
// )

// // minikatsuは、起床時間の登録等のモニカツの会話ロジックを司る。

// // セッションにuseridを登録して、モニカツ予約フラグを立てる。
// func (c *server.LineConn) ResisterMonikatsu(message *linebot.TextMessage) {
// 	// kvsマイクロサービスにセッションをinsert
// 	err := setContext(c.event.Source.UserID, "2.1")
// 	if err != nil {
// 		log.Print(err)
// 	}

// 	resp := linebot.NewTextMessage(constant.MONIKATSU_RESISTER)
// 	_, err = c.bot.ReplyMessage(c.event.ReplyToken, resp).Do()
// 	if err != nil {
// 		log.Print(err)
// 	}
// }
