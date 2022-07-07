package string_test

import (
	"strconv"
	"strings"
	"testing"
)

func TestStringFn(t *testing.T) {
	s := "A,B,C,D,E,F,G,H"
	// []rune
	parts := strings.Split(s, ",")
	for _, v := range parts {
		t.Log(v)
	}
	t.Log(strings.Join(parts, "-"))
}

func TestConvert(t *testing.T) {
	s := strconv.Itoa(10)
	t.Log("str" + s)
	if v, err := strconv.Atoi("10"); err == nil {
		t.Log(v)
	}
}
