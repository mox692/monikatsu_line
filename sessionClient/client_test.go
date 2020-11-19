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
		if status.StatusCode != 200 {
			t.Errorf("setstatuscode  is not collect: want %d, got %d\n", 200, status.StatusCode)
		}
	}

	expect := map[string]string{
		"etw":  "2.1",
		"teqw": "2.2",
		"fds":  "0",
	}

	for key, _ := range cases {
		status, err := GetContext(key)

		if err != nil {
			t.Errorf("err occur: %s", err.Error())
		}

		if status.StatusCode != 200 {
			t.Errorf("err happen: statuscode is %d, not %d\n", status.StatusCode, 200)
		}

		if status.Data != expect[key] {
			t.Errorf("status.Data get %s, not %s\n", status.Data, expect[key])
		}
	}
}
