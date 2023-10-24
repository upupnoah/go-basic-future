package slice

import (
	"reflect"
	"testing"
)

func TestDelete(t *testing.T) {
	type args[T any] struct {
		src   []T
		index int
	}
	type testCase[T any] struct {
		name    string
		args    args[T]
		want    []T
		want1   T
		wantErr bool
	}
	tests := []testCase[int]{
		{
			name:    "index out of range",
			args:    args[int]{src: []int{1, 2, 3}, index: 12},
			want:    nil,
			want1:   0,
			wantErr: true,
		},
		{
			name:    "index less than 0",
			args:    args[int]{src: []int{1, 2, 3}, index: -1},
			want:    nil,
			want1:   0,
			wantErr: true,
		},
		{
			name:    "index last",
			args:    args[int]{src: []int{1, 2, 3}, index: 2},
			want:    []int{1, 2},
			want1:   3,
			wantErr: false,
		},
		{
			name:    "index middle",
			args:    args[int]{src: []int{1, 2, 3}, index: 1},
			want:    []int{1, 3},
			want1:   2,
			wantErr: false,
		},
		{
			name:    "index 0",
			args:    args[int]{src: []int{1, 2, 3}, index: 0},
			want:    []int{2, 3},
			want1:   1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := Delete(tt.args.src, tt.args.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Delete() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Delete() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
