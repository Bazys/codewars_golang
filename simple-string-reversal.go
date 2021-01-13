package kata

func Solve(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i++ {
		for ; i < j && r[i] == ' '; i++ {
		}
		for ; i < j && r[j] == ' '; j-- {
		}
		r[i], r[j] = r[j], r[i]
		j--
	}
	return string(r)
}
