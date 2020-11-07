package constant

import (
	"fmt"
	"time"
)

var (
	HelpMessage = ""

	MONIKATSU_RESISTER = `明日の予約な！了解明日は何時に起きるんだ??
	「7:30」「0:00」「13:15」みたいな形で入力してくれ」
	やっぱやめる場合は「中止」って打ち込んでくれ`

	MONIKATSU_WAKEUP_RESISTER = func(times string) string {
		date := time.Now()
		datestr := fmt.Sprintf("%s/%s", date.Month(), date.Day())
		return fmt.Sprintf("了解！明日の%s日の%s時に起床するんだな？よかったらyes,違ったらnoを入力してくれ\n", datestr, times)
	}

	MONIKATSU_WRONGTIME_INPUT    = `入力された時刻が正しくないな。。。`
	MONIKATSU_CONFIRM_WAKEUPTIME = `入力された時刻が正しくないな。。。`
)