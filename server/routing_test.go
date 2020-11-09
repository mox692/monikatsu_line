package server

import (
	"fmt"
	"testing"

	"github.com/line/line-bot-sdk-go/linebot"
)

func Test_parseInputTime(t *testing.T) {
	input := []string{
		"5:50",
		"0:23",
		"13:20",
		"19:67",
		"25:20",
	}

	expect := []string{
		"5:50",
		"0:23",
		"13:20",
		"",
		"",
	}

	for i, v := range input {
		get, _ := parseInputTime(v)
		fmt.Printf("result: %s\n", get)
		if get != expect[i] {
			t.Errorf("want %s, but get %s\n", expect[i], get)
		}
	}
}

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
