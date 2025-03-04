package random

import (
	"testing"
)

func TestRandom_Intn(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "generate random number between 0 and 10",
			args: args{n: 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRandom()
			got := r.Intn(tt.args.n)
			if got < 0 || got >= tt.args.n {
				t.Errorf("Random.Intn() = %v, want %v", got, tt.args.n)
			}
		})
	}
}
