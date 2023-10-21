package generic

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	type args[T Number] struct {
		vals []T
	}
	type testInt[T Number] struct {
		name string
		args args[T]
		want T
	}
	tests := []testInt[int]{
		{
			name: "Test case 1",
			args: args[int]{vals: []int{1, 2, 3}},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.args.vals...); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOthers(t *testing.T) {
	fmt.Println(Sum[int](1, 2, 3))
	fmt.Println(Sum[float32](1.1, 2.2, 3.3))
	fmt.Println(Sum[float64](1.1, 2.2, 3.3))
	fmt.Println(Sum[int64](1, 2, 3))
	fmt.Println(Sum[Integer](1, 2, 3))
}
