package server

import "testing"

func Test_parseInputTime(t *testing.T) {

	m := monikatsu{}

	input := map[string]bool{
		"12:30": true,
		"7:15":  true,
		"6:09":  true,
		"0:00":  true,  // OK
		"00:00": false, // NG
		"15:00": true,
		"24:00": false, // NG
		"25:00": false,
		"5:60":  false,
	}

	for k, v := range input {

		time, err := m.parseInputTime(k)
		if v && err != nil {
			t.Errorf("parseInputTime err, value is %s, expect %v, got err, %s", k, v, err.Error())
		}
		if !v && err == nil {
			t.Errorf("parseInputTime err, expect %T, got NoErr ", v)
		}
		if v && time != k {
			t.Errorf("parse err: expect %s, got %s", k, time)
		}
	}
}
