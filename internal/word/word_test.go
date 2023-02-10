package word

import (
	"testing"
)

func TestDifficulty(t *testing.T) {
	type args struct {
		w string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{name: "empty", args: args{""}, want: 0, wantErr: true},
		{name: "char", args: args{"a"}, want: 0, wantErr: true},
		{name: "article", args: args{"an"}, want: 0, wantErr: true},
		{name: "axe word", args: args{"axe"}, want: 5, wantErr: false},
		{name: "toy word", args: args{"toy"}, want: 7, wantErr: false},
		{name: "deer word", args: args{"deer"}, want: 2, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Difficulty(tt.args.w)
			if (err != nil) != tt.wantErr {
				t.Errorf("Difficulty() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Difficulty() got = %v, want %v", got, tt.want)
			}
		})
	}
}
