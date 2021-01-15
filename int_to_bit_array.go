package kata

func IntToBitArray(n uint64) (res []uint8) {
	var i uint8 = 0
	for n > 0 {
		if (n & 1) != 0 {
			res = append(res, i)
		}
		i++
		n >>= 1
	}
	return res
}

func BitArrayToInt(arr []uint8) uint64 {
	var res uint64
	for _, el := range arr {
		res |= (1 << el)
	}
	return res
}
