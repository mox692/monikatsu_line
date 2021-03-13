package restClient

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/mox692/monikatsu_line/session/rest/server/types"
)




func SetSession() *types.SetSessionResponse {
	url := "https://qiita.com/api/v2/authenticated_user/items?page=1&per_page=20"

	// var body io.Reader
	req, err := http.NewRequest("POST", url, nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		os.Exit(1)
	}
	restSetSessionResponse := &types.SetSessionResponse{}
	json.NewDecoder(resp.Body).Decode(&restSetSessionResponse)
	return restSetSessionResponse
}

func GetSession() *types.GetSessionResponse {
	url := "https://qiita.com/api/v2/authenticated_user/items?page=1&per_page=20"

	// var body io.Reader
	req, err := http.NewRequest("GET", url, nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		os.Exit(1)
	}
	restGetSessionResponse := &types.GetSessionResponse{}
	json.NewDecoder(resp.Body).Decode(&restGetSessionResponse)
	return restGetSessionResponse
}