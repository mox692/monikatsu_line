package server

import (
	"database/sql"
	"monikatsuline/constant"
	"monikatsuline/database/repository"

	"github.com/line/line-bot-sdk-go/linebot"
)

type resister struct {
	LineConn   LineConn
	statusCode int
}

// userが投げたapnameがdb内に存在するかを
func (r *resister) askAppName(message *linebot.TextMessage) error {

	_, err := repository.SelectByUserName(message.Text)

	if err == sql.ErrNoRows {
		err = r.resisterReply(constant.RESISTER_USERNAME_NOTFOUND)
	}
	if err != nil {
		return err
	}

	err = setContext(r.LineConn.event.Source.UserID, "1.2")
	if err != nil {
		return err
	}

	r.resisterReply(constant.RESISTER_ASK_PASSWORD)
	if err != nil {
		return err
	}
	return nil
}

// password認証。認証されたらLineUserをinsert
func (r *resister) askPassword() error {
	return nil
}

func (r *resister) resisterReply(text string) error {
	resp := linebot.NewTextMessage(constant.RESISTER_USERNAME_NOTFOUND)
	_, err := r.LineConn.bot.ReplyMessage(r.LineConn.event.ReplyToken, resp).Do()
	if err != nil {
		return err
	}
	return nil
}
