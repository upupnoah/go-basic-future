package generic

import (
	"reflect"
	"testing"
)

// 表格驱动测试（Table-Driven Testing）

func TestSwap(t *testing.T) {
	type args[T any] struct {
		p Pair[any]
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want Pair[T]
	}
	tests := []testCase[any]{ // 指定具体的类型为 int
		{
			name: "Test case 1",
			args: args[any]{p: Pair[any]{A: 1, B: 2}},
			want: Pair[any]{A: 2, B: 1},
		},
		{
			name: "Test case 2",
			args: args[any]{p: Pair[any]{A: 5, B: 4}},
			want: Pair[any]{A: 4, B: 5},
		},
		// 你可以添加更多的测试用例
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//t.Logf("%+v  %+v", Swap(tt.args.p), tt.want)
			if got := Swap(tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Swap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyType(t *testing.T) {
	person := Person{name: "Noah"}
	myValue := MyType[Person]{value: person}
	t.Logf("%+v", myValue)
}
