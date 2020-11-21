package sessionClient

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/mox692/monikatsu_line/session"

	"google.golang.org/grpc"
)

var serverAddr string = "localhost:9090"
var opts []grpc.DialOption

// SetContext はuserIDとsessioncodeを用いてセッションを登録します。
func SetContext(userID string, sessionCode string) (*session.SetStatus, error) {
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		log.Fatal("err: %w", err)
	}

	client := session.NewSessionClient(conn)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	setStatus, err := client.SetSession(ctx, &session.SessionRequest{UserID: userID, StatusID: sessionCode})
	if err != nil {
		fmt.Printf("err : %s\n", err)
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
