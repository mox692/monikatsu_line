package grpcClient

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/mox692/monikatsu_line/session/grpc/server/session"
	"google.golang.org/grpc"
)

// var serverAddr string = os.Getenv("SERVER_HOST") + os.Getenv("SERVER_PORT")
var serverAddr string = os.Getenv("GRPC_SERVER_HOST_CONTAINER") + os.Getenv("GRPC_SERVER_PORT")
var opts []grpc.DialOption

// SetContext はuserIDとsessioncodeを用いてセッションを登録します。
func SetContext(userID string, sessionCode string) (*session.SetStatus, error) {
	log.Println("called in SetContext")
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		log.Printf("serveraddr: %s\n", serverAddr)
		log.Fatal("err: %w", err)
	}
	log.Printf("conn success!!\n")
	client := session.NewSessionClient(conn)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	setStatus, err := client.SetSession(ctx, &session.SessionRequest{UserID: userID, StatusID: sessionCode})
	if err != nil {
		log.Printf("client.SetSession : %s\n", err)
		log.Printf("serveraddr: %s\n", serverAddr)
		return &session.SetStatus{}, err
	}
	fmt.Printf("request success!\n status: %+v\n", setStatus)
	return setStatus, nil
}

// GetContext はuserIDから、そのuserの会話におけるコンテキストを取得します。
// 内部でgRPCのGetSessionメソッドを使用しています。
func GetContext(userID string) (*session.GetStatus, error) {
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		log.Fatal("err: %w", err)
	}

	client := session.NewSessionClient(conn)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	getStatus, err := client.GetSession(ctx, &session.SessionRequest{UserID: userID})
	if err != nil {
		fmt.Printf("err : %s\n", err)
		return &session.GetStatus{}, err
	}

	fmt.Printf("request success!\n status: %+v\n", getStatus)
	return getStatus, nil
}
