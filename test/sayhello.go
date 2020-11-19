package test

import (
	"log"
	"net/http"
)

func Sayhello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world from monikatsu server!!"))
	log.Println("get access")
}
