package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_asChan(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "nil arguments",
			args: args{nums: nil},
			want: []int{},
		},
		{
			name: "empty incomming array",
			args: args{nums: make([]int, 0)},
			want: []int{},
		},
		{
			name: "nil arguments",
			args: args{nums: []int{1, 2, 3}},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			gotCh := asChan(tt.args.nums...)
			got := make([]int, 0, len(tt.args.nums))
			for val := range gotCh {
				got = append(got, val)
			}
			require.Equal(t, tt.want, got)
		})
	}
}

func Test_sq(t *testing.T) {
	type args struct {
		in func(t *testing.T) <-chan int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "closed chan",
			args: args{
				in: func(t *testing.T) <-chan int {
					ch := make(chan int)
					defer close(ch)
					return ch
				},
			},
			want: []int{},
		},
		{
			name: "chan with one value",
			args: args{
				in: func(t *testing.T) <-chan int {
					return asChan(1)
				},
			},
			want: []int{1},
		},
		{
			name: "chan with many values",
			args: args{
				in: func(t *testing.T) <-chan int {
					return asChan(1, 2, 3)
				},
			},
			want: []int{1, 4, 9},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			gotCh := sq(tt.args.in(t))
			got := []int{}
			for val := range gotCh {
				got = append(got, val)
			}
			require.Equal(t, tt.want, got)
		})
	}
}
