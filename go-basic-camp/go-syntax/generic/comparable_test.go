package generic

import (
	"cmp"
	"reflect"
	"testing"
)

func TestMax(t *testing.T) {
	type args[T cmp.Ordered] struct {
		x T
		y T
	}
	type testCase[T cmp.Ordered] struct {
		name string
		args args[T]
		want T
	}
	testsInt := []testCase[int]{
		{
			name: "Test case 1",
			args: args[int]{x: 1, y: 2},
			want: 2,
		},
		{
			name: "Test case 2",
			args: args[int]{x: 5, y: 4},
			want: 5,
		},
	}
	testsFloat := []testCase[float64]{
		{
			name: "Test case 1",
			args: args[float64]{x: 1.1, y: 2.2},
			want: 2.2,
		},
		{
			name: "Test case 2",
			args: args[float64]{x: 5.5, y: 4.4},
			want: 5.5,
		},
	}
	testsString := []testCase[string]{
		{
			name: "Test case 1",
			args: args[string]{x: "a", y: "b"},
			want: "b",
		},
		{
			name: "Test case 2",
			args: args[string]{x: "apple", y: "banana"},
			want: "banana",
		},
	}
	for _, tt := range testsInt {
		t.Run(tt.name, func(t *testing.T) {
			if got := Max(tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Max() = %v, want %v", got, tt.want)
			}
		})
	}
	for _, tt := range testsFloat {
		t.Run(tt.name, func(t *testing.T) {
			if got := Max(tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Max() = %v, want %v", got, tt.want)
			}
		})
	}
	for _, tt := range testsString {
		t.Run(tt.name, func(t *testing.T) {
			if got := Max(tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Max() = %v, want %v", got, tt.want)
			}
		})
	}
}
