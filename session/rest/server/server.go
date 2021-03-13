package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/mox692/monikatsu_line/session/kvs"
	"github.com/mox692/monikatsu_line/session/rest/server/types"
	"golang.org/x/xerrors"
)

func main(){
	serverPort := ":" + os.Getenv("SESSION_REST_SERVER_PORT")
	s := http.Server{Addr: serverPort}

	// TODO: redisとの疎通

	// 以下はテスト用
	// http.HandleFunc("/hello", test.Sayhello)
	// http.HandleFunc("/insert_line_user", test.InsertLineUser)
	// http.HandleFunc("/grpc_test", test.ConnGRPC)
	// http.HandleFunc("/setcontext", test.SetSessionTest)
	// http.HandleFunc("/getcontext", test.GetSessionTest)
	http.HandleFunc("/setcontext", setSession)
	http.HandleFunc("/getcontext", getSession)

	go func(){
		log.Printf("server is running on port %s",serverPort)
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

// redisに対してsetを実行
func setSession(w http.ResponseWriter, r *http.Request) {
	
	var setSessionRequest = &types.SessionRequest{}
	err := json.NewDecoder(r.Body).Decode(setSessionRequest)
	log.Printf("get reqest!!\n")
	
	err = kvs.Set(setSessionRequest.UserID, setSessionRequest.StatusID, kvs.Conn)
	if err != nil {
		serverErr(w,&types.SetSessionResponse{StatusCode: 500, ErrMessage: err.Error()} ,xerrors.Errorf("kvs.Set err :%w", err))
	}
	success(w,&types.SetSessionResponse{StatusCode: 200} )
}

func getSession(w http.ResponseWriter, r *http.Request){
	var sessionRequest = &types.SessionRequest{}
	err := json.NewDecoder(r.Body).Decode(sessionRequest)
	log.Printf("get reqest!!\n")
	
	status, err := kvs.Get(sessionRequest.UserID, kvs.Conn)
	if err != nil {
		serverErr(w,&types.GetSessionResponse{StatusCode: 500, Data: "", ErrMessage: err.Error()} ,xerrors.Errorf("kvs.Set err :%w", err))
	}
	success(w,&types.GetSessionResponse{StatusCode: 200, Data: status})
}

// structをjsonに変換してwriterに書き込む
func success(w http.ResponseWriter, response interface{}){	
	jsonData, err := json.Marshal(response)
	if err !=nil {
		serverErr(w,&types.SetSessionResponse{StatusCode: 500, ErrMessage: err.Error()},err)
	}
	w.Write(jsonData)
	return 
}

// writerにerror structを書き込む、error logを残す
func serverErr(writer http.ResponseWriter, response interface{}, err error){
	jsonData, err := json.Marshal(response)
	err = xerrors.Errorf(": %w", err)
	log.Printf("%+v\n", err)
	writer.Write(jsonData)
	return
}

