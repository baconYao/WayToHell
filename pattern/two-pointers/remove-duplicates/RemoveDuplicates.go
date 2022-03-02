package removeduplicates

// time: O(n) | space: O(1)
func RemoveDuplicates(input []int) int {
	frontPointer := 0
	distinctCount := 0

	if len(input) > 0 {
		distinctCount = 1
	}

	for rearPointer, rValue := range input {
		if frontPointer == rearPointer {
			rearPointer += 1
			continue
		}
		fValue := input[frontPointer]
		if fValue != rValue {
			distinctCount += 1
			frontPointer = rearPointer
		}
	}

	return distinctCount
}
