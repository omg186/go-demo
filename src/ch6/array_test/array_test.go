package array_test

import "testing"

func TestArrayInit(t *testing.T) {
	var arr [3]int
	arr1 := [4]int{1, 2, 3, 4}
	arr3 := [...]int{1, 3, 4, 5}
	arr1[1] = 5
	t.Log(arr[1], arr[2])
	t.Log(arr1, arr3)
	// if arr1 == arr3 {
	t.Log(arr1 == arr3)
	// }
}

func TestArrayRange(t *testing.T) {
	arr1 := [4]int{1, 2, 3, 4}
	for index, e := range arr1 {
		t.Log(index, e)
	}
}

func TestArraySection(t *testing.T) {
	arr3 := [...]int{1, 2, 3, 4, 5}

	arr3_sec := arr3[1:2]
	// t.Log(arr3_sec, arr3)
	t.Log(len(arr3_sec), cap(arr3_sec))
	// t.Log(len(arr3), cap(arr3))
	// arr3_sec = append(arr3_sec, 10)
	// t.Log(len(arr3_sec), cap(arr3_sec))
	// arr3_sec = append(arr3_sec, 11)
	// t.Log(len(arr3_sec), cap(arr3_sec))
	// arr3_sec[1] = 20
	// t.Log(arr3_sec, arr3)
}
