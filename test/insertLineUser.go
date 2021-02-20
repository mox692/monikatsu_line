package test

import (
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/mox692/monikatsu_line/database/repository"
)

// /insert_line_userへのアクセスで、LineUserのinsertをテストします。
func InsertLineUser(w http.ResponseWriter, r *http.Request) {

	uuidObj, err := uuid.NewUUID()
	if err != nil {
		fmt.Printf("err: %s\n", err.Error())
		w.Write([]byte("errだああ"))
		return
	}

	// lineuserの作成
	u := repository.LineUser{UserID: uuidObj.String(), LineUserToken: uuidObj.String()}
	err = u.InsertLineUser()

	if err != nil {
		fmt.Printf("err: %s\n", err.Error())
		w.Write([]byte("errだああ"))
		return
	}
	log.Println("get access!")
	w.Write([]byte("ok"))
}
