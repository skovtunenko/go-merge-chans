package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_merge(t *testing.T) {
	type args struct {
		ch1 <-chan int
		ch2 <-chan int
	}
	tests := []struct {
		name string
		args func(t *testing.T) args
		want []int
	}{
		{
			name: "empty chans",
			args: func(t *testing.T) args {
				t.Helper()
				return args{
					ch1: asChan(),
					ch2: asChan(),
				}
			},
			want: []int{},
		},
		{
			name: "empty first chan",
			args: func(t *testing.T) args {
				t.Helper()
				return args{
					ch1: asChan(),
					ch2: asChan(1, 2, 3),
				}
			},
			want: []int{1, 2, 3},
		},
		{
			name: "empty second chan",
			args: func(t *testing.T) args {
				t.Helper()
				return args{
					ch1: asChan(1, 2, 3),
					ch2: asChan(),
				}
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			args := tt.args(t)
			gotCh := mergeTwo(args.ch1, args.ch2)
			got := []int{}
			for val := range gotCh {
				got = append(got, val)
			}
			require.Equal(t, tt.want, got)
		})
	}
}
