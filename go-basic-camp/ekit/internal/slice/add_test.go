package slice

import (
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {
	type args[T any] struct {
		src     []T
		element T
		index   int
	}
	type testCase[T any] struct {
		name    string
		args    args[T]
		want    []T
		wantErr bool
	}
	tests := []testCase[int]{
		{
			name: "index middle",
			args: args[int]{src: []int{1, 2, 3}, element: 4, index: 1},
			want: []int{1, 4, 2, 3},
		},
		{
			name: "index 0",
			args: args[int]{src: []int{1, 2, 3}, element: 4, index: 0},
			want: []int{4, 1, 2, 3},
		},
		{
			name:    "index out of range",
			args:    args[int]{src: []int{1, 2, 3}, element: 4, index: 12},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "index less than 0",
			args:    args[int]{src: []int{1, 2, 3}, element: 4, index: -1},
			want:    nil,
			wantErr: true,
		},
		{
			name: "index last",
			args: args[int]{src: []int{1, 2, 3}, element: 4, index: 3},
			want: []int{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Add(tt.args.src, tt.args.element, tt.args.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Add() got = %v, want %v", got, tt.want)
			}
		})
	}
}
