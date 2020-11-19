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

// userが投げたapnameがDBに存在するか確認する関数です。
func (r *resister) askAppName(message *linebot.TextMessage) error {

	_, err := repository.SelectUserByUserName(message.Text)

	if err == sql.ErrNoRows {
		err = r.resisterReply(constant.RESISTER_USERNAME_NOTFOUND)
		if err != nil {
			return err
		}
		err = setContext(r.LineConn.event.Source.UserID, "0.0")
		if err != nil {
			return err
		}
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

// userが投げたpasswordがDBにあルカ確認する関数。認証されたらLineUserをinsert
func (r *resister) askPassword(message *linebot.TextMessage) error {
	_, err := repository.SelectUserByPassword(message.Text)

	if err == sql.ErrNoRows {
		err = r.resisterReply(constant.RESISTER_PASSWORD_NOTFOUND)
		if err != nil {
			return err
		}
		err = setContext(r.LineConn.event.Source.UserID, "0")
		if err != nil {
			return err
		}
	}
	if err != nil {
		return err
	}

	err = setContext(r.LineConn.event.Source.UserID, "0")
	if err != nil {
		return err
	}

	err = r.insertLineUser()
	if err != nil {
		return err
	}

	r.resisterReply(constant.RESISTER_OK)
	if err != nil {
		return err
	}
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

func (r *resister) insertLineUser() error {
	return nil
}
