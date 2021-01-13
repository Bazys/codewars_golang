package kata

import "sort"

// Given an array/list [] of n integers , Arrange them in a way similar to the to-and-fro movement of a Pendulum
// * The Smallest element of the list of integers , must come in center position of array/list.
// * The Higher than smallest , goes to the right .
// * The Next higher number goes to the left of minimum number and So on , in a to-and-fro manner similar to that of a Pendulum.
// Notes
// * Array/list size is at least 3.
// * In Even array size , The minimum element should be moved to (n-1)/2 index (considering that indexes start from 0)
// * Repetition of numbers in the array/list could occur , So (duplications are included when Arranging).
// Explanation:
// pendulum ([6, 6, 8 ,5 ,10]) ==> [10, 6, 5, 6, 8]
// * Since , 5 is the The Smallest element of the list of integers , came in The center position of array/list
// * The Higher than smallest is 6 goes to the right of 5.
// * The Next higher number goes to the left of minimum number and So on .
// * Remember , Duplications are included when Arranging , Don't Delete Them .
func Pendulum(values []int) []int {
	sort.Ints(values)

	// too clever version
	// l, r := len(values), make([]int, len(values))
	// for i, v := range values {
	// 	r[l/2-i/2+i*(i&1)-((l+1)&1)] = v
	// }
	// return r

	var dst []int
	for i, v := range values {
		if i%2 == 0 {
			dst = append([]int{v}, dst...)
		} else {
			dst = append(dst, v)
		}
	}
	return dst
}
