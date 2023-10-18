package main

import "testing"

func TestMyPow(t *testing.T) {
	var powTests = []struct {
		in       int
		expected int
	}{
		{2, 4},
		{3, 9},
		{10, 100},
		{9, 81},
	}
	for _, i := range powTests {
		temp := i.in
		t1 := i.in
		MyPow(&t1)
		if i.expected != t1 {
			t.Errorf("MyPow(%d) = %d; expcted %d", temp, t1, i.expected)
		}
	}
}
