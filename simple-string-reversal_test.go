package kata

import "testing"

func TestSolve(t *testing.T) {
	tests := []struct {
		str  string
		want string
	}{
		{str: "", want: ""},
		{str: "codewars", want: "srawedoc"},
		{str: "your code", want: "edoc ruoy"},
		{str: "your code rocks", want: "skco redo cruoy"},
	}
	for _, tt := range tests {
		if got := Solve(tt.str); got != tt.want {
			t.Errorf("Solve() = %v, want %v", got, tt.want)
		}
	}
}
