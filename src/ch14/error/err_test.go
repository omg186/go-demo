package error

import (
	"errors"
	"testing"
)

func getFibonacci(n int) ([]int, error) {
	if n < 2 || n > 100 {
		return nil, errors.New("n should be not 2 ~ 100")
	}
	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}

func TestGetFibonacci(t *testing.T) {
	t.Log(getFibonacci(1))
}
