package func_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMultipleValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}
func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}
func TestFn(t *testing.T) {
	i, i2 := returnMultipleValues()
	t.Log(i, i2)
	s := timeSpent(func(op int) int {
		time.Sleep(time.Second * 1)
		return op
	})(10)
	t.Log(s)
}

func sum(ops ...int) int {

	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

func TestVarParam(t *testing.T) {
	t.Log(sum(1, 2, 3))
	t.Log(sum(1, 2, 3, 4))
}

// 延迟函数
func TestDefer(t *testing.T) {
	defer func() {
		t.Log("Clear resources")
	}()
	t.Log("Start")
	panic("Fatal error")
}
