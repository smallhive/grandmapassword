package keyboard

import (
	"testing"
)

func TestDistance(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "fh", args: args{"f", "h"}, want: 2},
		{name: "ae", args: args{"a", "e"}, want: 3},
		{name: "qq", args: args{"q", "q"}, want: 0},
		{name: "qw", args: args{"q", "w"}, want: 1},
		{name: "qa", args: args{"q", "a"}, want: 1},
		{name: "qz", args: args{"q", "z"}, want: 2},
		{name: "qp", args: args{"q", "p"}, want: 9},
		{name: "ql", args: args{"q", "l"}, want: 9},
		{name: "qm", args: args{"q", "m"}, want: 8},
		{name: "ta", args: args{"t", "a"}, want: 5},
		{name: "tg", args: args{"t", "g"}, want: 1},
		{name: "tl", args: args{"t", "l"}, want: 5},
		{name: "bq", args: args{"b", "q"}, want: 6},
		{name: "bt", args: args{"b", "t"}, want: 2},
		{name: "bp", args: args{"b", "p"}, want: 7},
		{name: "ba", args: args{"b", "a"}, want: 5},
		{name: "bh", args: args{"b", "h"}, want: 2},
		{name: "bl", args: args{"b", "l"}, want: 5},
		{name: "bz", args: args{"b", "z"}, want: 4},
		{name: "bm", args: args{"b", "m"}, want: 2},
		{name: "bb", args: args{"b", "b"}, want: 0},
		{name: "bb", args: args{"p", "a"}, want: 10},
		{name: "pa", args: args{"p", "a"}, want: 10},
		{name: "ps", args: args{"p", "s"}, want: 9},
		{name: "pm", args: args{"p", "m"}, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Distance([]rune(tt.args.a)[0], []rune(tt.args.b)[0]); got != tt.want {
				t.Errorf("Distance() = %v, want %v", got, tt.want)
			}
		})
	}
}
