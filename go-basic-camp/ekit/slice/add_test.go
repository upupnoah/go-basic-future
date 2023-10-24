package slice

import (
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {
	type args[Src any] struct {
		src     []Src
		element Src
		index   int
	}
	type testCase[Src any] struct {
		name    string
		args    args[Src]
		want    []Src
		wantErr bool
	}
	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
		// TODO: Add test cases.
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
