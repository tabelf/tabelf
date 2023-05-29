package utils

import "strconv"

func MaxInt(arr ...int) int {
	max := arr[0]
	for _, i := range arr {
		if i > max {
			max = i
		}
	}
	return max
}

func MinInt(arr ...int) int {
	min := arr[0]
	for _, i := range arr {
		if i < min {
			min = i
		}
	}
	return min
}

func IsNum(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}
