package sessionClient

import (
	"testing"
)

func Test_SetContext(t *testing.T) {

	cases := map[string]string{
		"etw":  "2.1",
		"teqw": "2.2",
		"fds":  "0",
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
