package kata

import (
	"strings"
)

func DuplicateEncode(word string) string {
	q := map[rune]int{}
	word = strings.ToLower(word)
	for _, el := range word {
		q[el]++
	}
	var res string
	for _, el := range word {
		if q[el] > 1 {
			res += ")"
		} else {
			res += "("

		}
	}
	return res
}
