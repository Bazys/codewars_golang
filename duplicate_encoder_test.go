package kata

import "testing"

func TestDuplicateEncode(t *testing.T) {
	tests := []struct {
		word string
		want string
	}{
		{word: "din", want: "((("},
		{word: "recede", want: "()()()"},
		{word: "Success", want: ")())())"},
		{word: "(( @", want: "))(("},
	}
	for _, tt := range tests {
		if got := DuplicateEncode(tt.word); got != tt.want {
			t.Errorf("DuplicateEncode() = %v, want %v", got, tt.want)
		}
	}
}
