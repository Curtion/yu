package utils

import "math"

func FormatMaxNum(num int64, byteLength int64) int64 {
	var res int64
	if num < 0 {
		res = 0
	} else if num > int64(math.Pow(2, float64(byteLength*8))-1) {
		res = int64(math.Pow(2, float64(byteLength*8)) - 1)
	} else {
		res = num
	}
	return res
}
