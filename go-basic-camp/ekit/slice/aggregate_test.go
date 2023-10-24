package slice

import (
	"github.com/upupnoah/go-basic-future/go-basic-camp/ekit"
	"testing"
)

func TestMax(t *testing.T) {
	type args[T ekit.RealNumber] struct {
		ts []T
	}
	type testCase[T ekit.RealNumber] struct {
		name string
		args args[T]
		want T
	}
	intTests := []testCase[int]{
		{
			name: "Max int 1",
			args: args[int]{ts: []int{1, 2, 3, 4, 5}},
			want: 5,
		},
		{
			name: "Max int 2",
			args: args[int]{ts: []int{5, 4, 3, 2, 1}},
			want: 5,
		},
		{
			name: "Max int 3",
			args: args[int]{ts: []int{1, 2, 3, 5, 4}},
			want: 5,
		},
	}
	float64Tests := []testCase[float64]{
		{
			name: "Max float64 1",
			args: args[float64]{ts: []float64{1.1, 2.2, 3.3, 4.4, 5.5}},
			want: 5.5,
		},
		{
			name: "Max float64 2",
			args: args[float64]{ts: []float64{5.5, 4.4, 3.3, 2.2, 1.1}},
			want: 5.5,
		},
		{
			name: "Max float64 3",
			args: args[float64]{ts: []float64{1.1, 2.2, 3.3, 5.5, 4.4}},
			want: 5.5,
		},
	}
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Max(tt.args.ts); got != tt.want {
				t.Errorf("Max() = %v, want %v", got, tt.want)
			}
		})
	}
	for _, tt := range float64Tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Max(tt.args.ts); got != tt.want {
				t.Errorf("Max() = %v, want %v", got, tt.want)
			}
		})
	}
}
