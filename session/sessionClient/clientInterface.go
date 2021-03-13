package sessionClient

import (
	"fmt"
	"log"
	"os"

	"github.com/mox692/monikatsu_line/session/grpc/grpcClient"
	"github.com/mox692/monikatsu_line/session/rest/restClient"
	"golang.org/x/xerrors"
)

const (
	REST = "REST"
	GRPC = "GRPC"
)

// 昔のsetsessionとかをここに大体する。
// 環境変数を読んで、restかgrpcかを切り替える

type SetStatus struct {
	StatusCode int32
	ErrMessage string
}

func SetSession(userID, sessionCode string)(*SetStatus, error) {

	clientType:= os.Getenv("SESSION_CLIENT_TYPE")
	log.Printf("client type is %s", clientType)
	switch clientType {
	case "REST":
		result := restClient.SetSession()
		return &SetStatus{StatusCode: result.StatusCode, ErrMessage: result.ErrMessage}, nil
	case "GRPC":
		setStatus, err := grpcClient.SetContext(userID, sessionCode)
		if err != nil {
			return nil, err
		}
		mock := &SetStatus{StatusCode: setStatus.StatusCode, ErrMessage: setStatus.ErrMessage}
		return mock, nil
	default:
		log.Printf("client type is %s", clientType)
		return nil, xerrors.New(fmt.Sprintf("env SESSION_CLIENT_TYPE is not correct, clientType: %s",clientType ))
	}
}

type GetStatus struct {
	StatusCode int32
	Data string
	ErrMessage string
}

func GetSession(userID string) (*GetStatus, error) {
	clientType:= os.Getenv("SESSION_CLIENT_TYPE")
	log.Printf("client type is %s", clientType)
	switch clientType {
	case "REST":
		result := restClient.GetSession()
		return &GetStatus{StatusCode: result.StatusCode, Data: result.Data, ErrMessage: result.ErrMessage} , nil
	case "GRPC":
		getStatus, err := grpcClient.GetContext(userID)
		if err != nil {
			return nil, err
		}
		mock := &GetStatus{StatusCode: getStatus.StatusCode, Data: getStatus.Data, ErrMessage: getStatus.ErrMessage}
		return mock, nil
	default:
		log.Printf("client type is %s", clientType)
		return nil, xerrors.New(fmt.Sprintf("env SESSION_CLIENT_TYPE is not correct, clientType: %s",clientType ))
	}
}