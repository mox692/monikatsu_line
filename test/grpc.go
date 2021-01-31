package test

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/mox692/monikatsu_line/session"
	"google.golang.org/grpc"
)

// var serverAddr string = os.Getenv("GRPC_SERVER_HOST") + ":" + os.Getenv("GRPC_SERVER_PORT")
var serverAddr string = "35.221.73.47:50051"
var opts []grpc.DialOption

func ConnGRPC(w http.ResponseWriter, r *http.Request) {
	log.Printf("gRPC addr: %s\n", serverAddr)
	ctx, _ := context.WithTimeout(context.Background(), time.Second*20)

	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		log.Printf("err: %w", err)
	}

	client := session.NewSessionClient(conn)

	res, err := client.ConnTest(ctx, &session.TestMessage{Msg: "access from monikatsu_servre!!!"})
	if err != nil {
		log.Printf("err: %w", err)
	}
	w.Write([]byte(res.Msg))
	log.Printf("success!! res: %s", res.Msg)
}
