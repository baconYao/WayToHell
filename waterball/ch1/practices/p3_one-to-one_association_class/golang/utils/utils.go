package utils

import "fmt"

func ShouldBeWithinRange(name string, num, min, max int) (int, error) {
	if min > max {
		return 0, fmt.Errorf("min must be <= Max")
	}
	if num < min || num > max {
		return 0, fmt.Errorf("%s must be within 0~100", name)
	}
	return num, nil
}
