package server

import (
	"testing"

	"github.com/line/line-bot-sdk-go/linebot"
)

func Test_judgeContext(t *testing.T) {

	inputs := []*LineConn{
		{
			event: &linebot.Event{
				Source: &linebot.EventSource{
					UserID: "motoyuki",
				},
			},
		},
		{
			event: &linebot.Event{
				Source: &linebot.EventSource{
					UserID: "yuki",
				},
			},
		},
	}

	messages := []*linebot.TextMessage{
		{
			Text: "hi",
		},
		{
			Text: "fdsa",
		},
	}

	for i, v := range inputs {
		err := v.judgeContext(messages[i])
		if err != nil {
			t.Errorf("err : %+v", err)
		}
	}

}
