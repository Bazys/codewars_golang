package kata

import (
	"strings"
)

func ReverseWords(s string) string {
	var res []string
	arr := strings.Split(s, " ")
	for i := 0; i < len(arr); i++ {
		word := []byte(arr[i])
		for j := len(word)/2 - 1; j >= 0; j-- {
			opp := len(word) - 1 - j
			word[j], word[opp] = word[opp], word[j]
		}
		res = append(res, string(word))
	}
	return strings.Join(res, " ")
}
