package groutine_test

import (
	"fmt"
	"testing"
	"time"
)

func TestGroutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Printf("i: %v\n", i)
		}(i)
	}
	time.Sleep(time.Second * 50)
}
