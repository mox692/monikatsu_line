package server

import (
	"fmt"
	"log"
	"regexp"

	"github.com/mox692/monikatsu_line/constant"

	"github.com/line/line-bot-sdk-go/linebot"
	"golang.org/x/xerrors"
)

type monikatsu struct {
	LineConn LineConn
}

func (m *monikatsu) setWakeupTime(message *linebot.TextMessage) {

	// 入力メッセージをvalidate
	time, err := m.parseInputTime(message.Text)
	if err != nil {
		resp := linebot.NewTextMessage(constant.MONIKATSU_WRONGTIME_INPUT)
		_, err = m.LineConn.bot.ReplyMessage(m.LineConn.event.ReplyToken, resp).Do()
		if err != nil {
			log.Print(err)
		}
		return
	}
	fmt.Println(time)

	// kvsマイクロサービスにセッションをinsert
	err = setContext(m.LineConn.event.Source.UserID, "2.2")
	if err != nil {
		log.Print(err)
	}

	resp := linebot.NewTextMessage(constant.MONIKATSU_WAKEUP_RESISTER(time))
	_, err = m.LineConn.bot.ReplyMessage(m.LineConn.event.ReplyToken, resp).Do()
	if err != nil {
		log.Print(err)
	}
}

func (m *monikatsu) parseInputTime(time string) (string, error) {
	r, _ := regexp.Compile("^([0-9]|1[0-9]|2[0-3]|):[0-5][0-9]")
	if !(r.Match([]byte(time))) {
		return "", xerrors.New("parse err!")
	}
	return time, nil
}
