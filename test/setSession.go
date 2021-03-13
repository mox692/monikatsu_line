package test

import (
	"log"
	"net/http"

	"github.com/mox692/monikatsu_line/session/grpc/grpcClient"
)

// /setcontext でのアクセスに対して、SetContextをテストします。。
func SetSessionTest(w http.ResponseWriter, r *http.Request) {
	log.Printf("called in SetSessionTest\n")
	status, err := grpcClient.SetContext("user", "2.3")
	if err != nil {
		log.Printf("err! %s", err.Error())
		w.Write([]byte("err!! check the server log"))
		return
	}
	log.Printf("success!! status: %d", status.StatusCode)
	w.Write([]byte("success"))
}

func GetSessionTest(w http.ResponseWriter, r *http.Request) {
	status, err := grpcClient.GetContext("user")

	if err != nil {
		log.Printf("err! %s", err.Error())
		w.Write([]byte("err!! check the server log"))
		return
	}
	log.Printf("success!! data: %s, statuscode: %d", status.Data, status.StatusCode)
	w.Write([]byte("success"))
}
