package test

import (
	"fmt"
	"log"
	"monikatsuline/database/repository"
	"net/http"
)

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
