package server

import (
	"fmt"
	"log"
	"regexp"
	"time"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/mox692/monikatsu_line/constant"
	"github.com/mox692/monikatsu_line/database/repository"
	"github.com/mox692/monikatsu_line/sessionClient"
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
	status, err := sessionClient.SetContext(m.LineConn.event.Source.UserID, "2.2")
	if err != nil {
		log.Print(err)
	}
	fmt.Printf("statusCode: %d", status.StatusCode)

	// TODO: 時刻確認stepを入れるのであれば、useridと時刻を保存するRPCを追加して
	//       ここで呼ぶようにする

	resp := linebot.NewTextMessage(constant.MONIKATSU_WAKEUP_RESISTER(time))
	_, err = m.LineConn.bot.ReplyMessage(m.LineConn.event.ReplyToken, resp).Do()
	if err != nil {
		log.Print(err)
	}
}

func (m *monikatsu) confirmWakeupTime(message *linebot.TextMessage) {

	if text := message.Text; text == "yes" || text == "YES" || text == "Yes" {
		// TODO: sessionで日付のdataを保存する？？
		t := time.Now()
		wd := &repository.WakeupData{"userid", m.LineConn.event.Source.UserID, "2/4", "7:00", timeToMonikatsuTime(t), timeToMonikatsuTime(t)}
		wd.Insert()
		resp := linebot.NewTextMessage(constant.MONIKATSU_WAKEUP_CONFIRM)
		_, err := m.LineConn.bot.ReplyMessage(m.LineConn.event.ReplyToken, resp).Do()
		if err != nil {
			log.Print(err)
		}
		err = setContext(m.LineConn.event.Source.UserID, "2.5")
		if err != nil {
			log.Print(err)
		}
	} else {
		resp := linebot.NewTextMessage(constant.MONIKATSU_CANCEL)
		_, err := m.LineConn.bot.ReplyMessage(m.LineConn.event.ReplyToken, resp).Do()
		if err != nil {
			log.Print(err)
		}
	}
}

func (m *monikatsu) checkWakeupTime(message *linebot.TextMessage) {
	// TODO: userIDを引数にして、db or Redisからwakeuodateの値を取ってくる
	wd := repository.CreateWakeupData()
	wakeupData, err := wd.SelectByUserID(m.LineConn.event.Source.UserID)
	if err != nil {
		log.Print(err)
	}

	resp := linebot.NewTextMessage(constant.CHECK_WAKEUP_TIME(wakeupData.WakeupTime))
	_, err = m.LineConn.bot.ReplyMessage(m.LineConn.event.ReplyToken, resp).Do()
	if err != nil {
		log.Print(err)
	}
}

func (m *monikatsu) unexpectedException() {
	resp := linebot.NewTextMessage(constant.UNEXPECTED_EXCEPTION)
	_, err := m.LineConn.bot.ReplyMessage(m.LineConn.event.ReplyToken, resp).Do()
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

func timeToMonikatsuTime(t time.Time) string {
	return fmt.Sprintf("%s-%s-%s %s:%s", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute())
}
