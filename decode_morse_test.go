package kata

import "testing"

func TestDecodeMorse(t *testing.T) {
	tests := []struct {
		morseCode string
		want      string
	}{
		{".... . -.--  .--- ..- -.. .", "HEY JUDE"},
		{".", "E"},
		{"- .... .  --.- ..- .. -.-. -.-  -... .-. --- .-- -.  ..-. --- -..-  .--- ..- -- .--. ...  --- ...- . .-.  - .... .  .-.. .- --.. -.--  -.. --- --. .-.-.-", "THE QUICK BROWN FOX JUMPS OVER THE LAZY DOG."},
	}
	for _, tt := range tests {
		if got := DecodeMorse(tt.morseCode); got != tt.want {
			t.Errorf("DecodeMorse() = %v, want %v", got, tt.want)
		}
	}
}
