package test

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mox692/monikatsu_line/database/repository"
)

// /insert_line_userへのアクセスで、LineUserのinsertをテストします。
func InsertLineUser(w http.ResponseWriter, r *http.Request) {

	// lineuserの作成
	u := repository.LineUser{UserID: "teast", LineUserToken: "testtoken"}
	err := u.InsertLineUser()

	if err != nil {
		fmt.Printf("err: %s\n", err.Error())
		w.Write([]byte("errだああ"))
		return
	}
	log.Println("get access!")
	w.Write([]byte("ok"))
}
