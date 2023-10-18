package main

import (
	"golang.org/x/tour/tree"
	"testing"
)

func TestWalk(t *testing.T) {
	type args struct {
		t  *tree.Tree
		ch chan int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"walk1",
			args{
				tree.New(1),
				make(chan int),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			go Walk(tt.args.t, tt.args.ch)
			for i := 1; i <= 10; i++ {
				if <-tt.args.ch != i {
					t.Errorf("Walk() = %v, want %v", <-tt.args.ch, i)
				}
			}
		})
	}
}

func TestSame(t *testing.T) {
	type args struct {
		t1 *tree.Tree
		t2 *tree.Tree
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"same1",
			args{
				tree.New(1),
				tree.New(1),
			},
			true,
		},
		{
			"same2",
			args{
				tree.New(1),
				tree.New(2),
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Same(tt.args.t1, tt.args.t2); got != tt.want {
				t.Errorf("Same() = %v, want %v", got, tt.want)
			}
		})
	}
}
