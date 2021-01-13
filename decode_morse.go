package kata

import "strings"

func DecodeMorse(morseCode string) string {
	morse := map[string]string{"a": ".-", "b": "-...", "c": "-.-.", "d": "-..", "e": ".", "f": "..-.",
		"g": "--.", "h": "....", "i": "..", "j": ".---", "k": "-.-", "l": ".-..", "m": "--", "n": "-.",
		"o": "---", "p": ".--.", "q": "--.-", "r": ".-.", "s": "...", "t": "-", "u": "..-", "v": "...-",
		"w": ".--", "x": "-..-", "y": "-.--", "z": "--..", ".": ".-.-.-", "0": "-----", "1": ".----", "2": "..---",
		"3": "...--", "4": "....-", "5": ".....", "6": "-....", "7": "--...", "8": "---..", "9": "----.",
	}
	revMorse := make(map[string]string)
	for k, v := range morse {
		revMorse[v] = k
	}
	var res string
	morseCode += " "
	var letter []byte
	max := 0
	count := len(morseCode)
	for index := 0; index < count; index++ {
		if morseCode[index] == ' ' {
			max++
			if max == 2 {
				res += " "
			} else if elem, ok := revMorse[string(letter)]; ok {
				res += elem
				letter = nil
			}
		} else {
			letter = append(letter, morseCode[index])
			max = 0
		}
	}
	return strings.ToUpper(res)
}
