package api

import "net/http"



func GetWakeupData(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("api test!"))

	// dbへの接続、get処理

	// 整形、return
}