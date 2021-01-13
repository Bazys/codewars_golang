package kata

import "testing"

func TestReverseWords(t *testing.T) {
	tests := []struct {
		str  string
		want string
	}{
		{str: "", want: ""},
		{str: "ab", want: "ba"},
		{str: "I want to be free", want: "I tnaw ot eb eerf"},
		{str: "ab", want: "ba"},
	}
	for _, tt := range tests {
		if got := ReverseWords(tt.str); got != tt.want {
			t.Errorf("reverseWords(\"%v\") = \"%v\", want \"%v\"", tt.str, got, tt.want)
		}
	}
}
