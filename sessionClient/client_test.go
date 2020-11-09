package sessionClient

import (
	"monikatsuline/server"
	"testing"
)

func Test_SetContext(t *testing.T) {

	cases := map[string]server.SessionCode{
		"etw":  server.MonikatsuFlag,
		"teqw": server.MonikatsuSetWakeupTime,
		"fds":  server.DefaultState,
	}

	for key, value := range cases {
		status, err := SetContext(key, value)
		if err != nil {
			t.Errorf("error occured: %v\n", err)
		}
		if status.SetStatusCode != "0" {
			t.Errorf("setstatuscode  is not collect: want %s, got %s\n", "0", status.SetStatusCode)
		}
	}
}
