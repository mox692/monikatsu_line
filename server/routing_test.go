package server

import (
	"fmt"
	"testing"
)

func Test_parseInputTime(t *testing.T) {
	input := []string{
		"5:50",
		"0:23",
		"13:20",
		"19:67",
		"25:20",
	}

	expect := []string{
		"5:50",
		"0:23",
		"13:20",
		"",
		"",
	}

	for i, v := range input {
		get, _ := parseInputTime(v)
		fmt.Printf("result: %s\n", get)
		if get != expect[i] {
			t.Errorf("want %s, but get %s\n", expect[i], get)
		}
	}
}
