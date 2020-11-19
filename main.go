package main

import (
	"context"
	"log"
	"monikatsuline/config"
	"monikatsuline/database"
	"monikatsuline/server"
	"monikatsuline/test"
	"net/http"
	"os"
	"os/signal"
	"time"
)

const channel_secret = "c8654b200557eb6744b21fc78f240d0a"
const channel_token = "G0k5a+PTDx8UU7s09xB3qD8viE7+AXGjrTJDBhiZ8Qs7j44nIKzlBgb0WFEqW7trLbJWcsX4HDXQJc3Bn+vJylb7S1sUnbzXiCWE9D4rUu9gsvYVVugCW4wc3dm6yQzlXsoPpBWIW4Kn1xZkOqRqjAdB04t89/1O/w1cDnyilFU="

func main() {

	// 環境変数のセット
	err := config.GetENV()
	if err != nil {
		log.Fatal(err)
	}

	// dbとの接続
	database.SetupDB()

	http.HandleFunc("/callback", server.JudgeEvent)

	// 以下はテスト用
	http.HandleFunc("/hello", test.Sayhello)
	http.HandleFunc("/insert_line_user", test.InsertLineUser)

	s := http.Server{Addr: ":8080"}

	go func() {
		log.Println("server runnning ...")
		if err := s.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	log.Printf("\nSIGNAL '%d' received, then shutting down...\n", <-quit)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Println(err)
	}
}
