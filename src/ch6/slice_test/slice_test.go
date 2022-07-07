package main

import "testing"

func TestSliceGrowing(t *testing.T) {
	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s))
	}
}

func TestSliceShareMemory(t *testing.T) {
	year := []string{"2013", "2014", "2015", "2016", "2017", "2018", "2019", "2020"}
	Q2 := year[2:5]
	t.Log(Q2, len(Q2), cap(Q2))
	sumer := year[4:7]
	t.Log(sumer, len(sumer), cap(sumer))
	sumer[0] = "Unknown"
	t.Log(Q2)
	t.Log(year)
}
func TestLiceComparing(t *testing.T) {
	// a := []int{1, 2, 3, 4}
	// b := []int{1, 2, 3, 4}
	// a.
	// if a == b {

	// }

	var x = make([]int, 0, 10)
	x = append(x, 1, 2, 3)
	t.Log(len(x), cap(x))
	y := append(x, 4)

	t.Logf("yyyyy----%d", y)
	z := append(x, 5)
	t.Log(x)
	t.Log(y)
	t.Log(z)

	t.Log("----------------------------------------------------------------")

	t.Log(x)
}
