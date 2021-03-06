package server

import (
	"database/sql"
	"log"

	"github.com/mox692/monikatsu_line/constant"

	"github.com/mox692/monikatsu_line/database/repository"
	"github.com/mox692/monikatsu_line/session/grpc/grpcClient"

	"golang.org/x/xerrors"

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
		status, err := grpcClient.SetContext(r.LineConn.event.Source.UserID, "0")
		if err != nil {
			return xerrors.Errorf("Err. message: %s, : %w", status.ErrMessage, err)
		}
		return nil
	}
	if err != nil {
		return err
	}

	status, err := grpcClient.SetContext(r.LineConn.event.Source.UserID, "1.2")
	if err != nil {
		return xerrors.Errorf("Err. message: %s, : %w", status.ErrMessage, err)
	}

	r.resisterReply(constant.RESISTER_ASK_PASSWORD)
	if err != nil {
		return err
	}
	return nil
}

// userが投げたpasswordがDBにあルカ確認する関数。認証されたらLineUserをinsert
func (r *resister) askPassword(message *linebot.TextMessage) error {
	user, err := repository.SelectUserByPassword(message.Text)

	if err == sql.ErrNoRows {
		err = r.resisterReply(constant.RESISTER_PASSWORD_NOTFOUND)
		if err != nil {
			return err
		}
		status, err := grpcClient.SetContext(r.LineConn.event.Source.UserID, "0")
		if err != nil {
			return xerrors.Errorf("Err. message: %s, : %w", status.ErrMessage, err)
		}
	}
	if err != nil {
		return err
	}

	status, err := grpcClient.SetContext(r.LineConn.event.Source.UserID, "0")
	if err != nil {
		return xerrors.Errorf("Err. message: %s, : %w", status.ErrMessage, err)
	}

	u := &repository.LineUser{UserID: user.ID, LineUserToken: r.LineConn.event.Source.UserID}
	err = u.InsertLineUser()
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

func (r *resister) unexpectedException() {
	resp := linebot.NewTextMessage(constant.UNEXPECTED_EXCEPTION)
	_, err := r.LineConn.bot.ReplyMessage(r.LineConn.event.ReplyToken, resp).Do()
	if err != nil {
		log.Print(err)
	}
}

func (r *resister) insertLineUser() error {
	return nil
}
