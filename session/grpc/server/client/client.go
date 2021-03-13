/* session serviceのclientの実装例です。 */

package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/mox692/monikatsu_line/session/grpc/server/session"
	"google.golang.org/grpc"
)

var serverAddr string = "localhost:9090"
var opts []grpc.DialOption

func main() {

	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		log.Fatal("err: %w", err)
	}

	client := session.NewSessionClient(conn)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	setStatus, err := client.SetSession(ctx, &session.SessionRequest{StatusID: "2.3", UserID: "fdaoisjgg"})
	if err != nil {
		fmt.Printf("err : %s\n", err)
		return
	}
	fmt.Printf("request success!\n status: %+v\n", setStatus)

	getStatus, err := client.GetSession(ctx, &session.SessionRequest{UserID: "fdaoisjsssgg"})
	if err != nil {
		fmt.Printf("err : %s\n", err)
		return
	}
	fmt.Printf("request success!\n getstatus: %+v\n", getStatus)

}
