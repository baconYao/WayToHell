package utils

func Sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func SquareInt(x int) int {
	return x * x
}
