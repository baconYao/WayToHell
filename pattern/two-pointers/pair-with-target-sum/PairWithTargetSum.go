package pairwithtargetsum

func PairWithTargetSum(target int, input []int) []int {
	// the pointer of start and end position
	sp := 0
	ep := len(input) - 1

	for sp < ep {
		sum := input[sp] + input[ep]
		if sum == target {
			return []int{sp, ep}
		} else if sum > target {
			ep -= 1
		} else {
			sp += 1
		}
	}

	return []int{-1, -1}
}
