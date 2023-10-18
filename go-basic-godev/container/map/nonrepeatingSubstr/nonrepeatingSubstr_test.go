package main

import (
	"testing"
)

func TestSubstr(t *testing.T) {
	tests := []struct {
		s   string
		ans int
	}{
		// Normal cases
		{"abcabcbb", 3},
		{"pwwkew", 3},

		// Edge cases
		{"", 0},
		{"b", 1},
		{"bbbbbbbb", 1},
		{"abcabcabcd", 4},
	}
	for _, tt := range tests {
		actual := lengthOfNonRepeatingSubstr(tt.s)
		//fmt.Println(actual)
		if actual != tt.ans {
			t.Errorf("got %d for input %s; "+
				"expected %d", actual, tt.s, tt.ans)
		}
	}
}

func BenchmarkSubstr(b *testing.B) {
	s, ans := "黑化黑灰化肥灰会挥发发灰黑讳为黑灰花会回飞", 8
	for i := 0; i < 13; i++ {
		s = s + s
	}

	b.Logf("len(s) = %d\n", len(s))
	b.ResetTimer() // 上面的时间不算

	for i := 0; i < b.N; i++ {
		actual := lengthOfNonRepeatingSubstr(s)
		if actual != ans {
			b.Errorf("got %d for input %s; "+
				"expected %d", actual, s, ans)
		}
	}
}
