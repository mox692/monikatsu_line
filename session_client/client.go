package session_client

import (
	"context"
	"fmt"
	"log"
	"time"

	"monikatsuline/server"
	"monikatsuline/session"

	"google.golang.org/grpc"
)

var serverAddr string = "localhost:9090"
var opts []grpc.DialOption

func SetContext(userID string, sessionCode server.SessionCode) (*session.SetStatus, error) {
	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		log.Fatal("err: %w", err)
	}

	client := session.NewSessionClient(conn)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	status, err := client.SetSession(ctx, &session.SessionRequest{UserID: userID, StatusID: string(sessionCode)})
	if err != nil {
		fmt.Printf("err : %s\n", err)
		return &session.SetStatus{}, err
	}
	fmt.Printf("request success!\n status: %+v\n", status)
	return status, nil
}
