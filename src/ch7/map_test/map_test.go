package map_test

import (
	"testing"
)

func TestInitMap(t *testing.T) {
	m1 := map[string]int{"one": 1, "two": 2, "three": 3}
	t.Log(m1["one"])
	t.Logf("len m1=%d", len(m1))
	m2 := map[string]int{}
	m2["a"] = 14
	t.Logf("len m2=%d", len(m2))
	t.Log(m2)
	m3 := make(map[string]int, 10)
	t.Logf("len m3=%d", len(m3))
}

func TestAccessNotExistingKey(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1])
	m1[2] = 0
	t.Log(m1[2])
	if v, ok := m1[3]; ok {
		t.Logf("Key 3 is value is %d", v)
	} else {
		t.Log("Key is not existing")
	}
}

func TestRangeMap(t *testing.T) {
	m1 := map[string]int{"one": 1, "two": 2, "three": 3}
	for key, v := range m1 {
		t.Log(key, v)
	}
}
