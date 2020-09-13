package map_ext

import "testing"

func TestMapWithFunValue(t *testing.T) {
	m1 := map[int]func(op int) int{}
	m1[1] = func(op int) int { return op }
	m1[2] = func(op int) int { return op * op }
	m1[3] = func(op int) int { return op * op * op }
	t.Log(m1[1](2), m1[2](2), m1[3](2))
}

func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{}
	mySet[1] = true
	n := 3
	if mySet[n] {
		t.Logf("%d is exiting", n)
	} else {
		t.Logf("%d is not exiting", n)
	}
	mySet[3] = true
	t.Log(len(mySet))
	delete(mySet, 1)
	n = 1
	if mySet[n] {
		t.Logf("%d is exiting", n)
	} else {
		t.Logf("%d is not exiting", n)
	}
}
