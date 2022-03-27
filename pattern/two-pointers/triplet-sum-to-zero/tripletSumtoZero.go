package tripletsumtozero

import (
	"sort"
)

// time: O(nlogn+N2) = O(N2) | space: O(N)
func TripletSumToZero(input []int) [][3]int {
	output := [][3]int{}
	sort.Ints(input)
	basePtr := 0
	frontPtr := basePtr + 1
	rearPtr := len(input) - 1

	for basePtr < len(input) {
		// Skip same element
		if basePtr > 0 && input[basePtr] == input[basePtr-1] {
			basePtr += 1
			continue
		}
		baseValue := input[basePtr]
		for frontPtr < rearPtr {
			frontValue := input[frontPtr]
			rearValue := input[rearPtr]
			if baseValue+frontValue+rearValue == 0 {
				output = append(output, [3]int{baseValue, frontValue, rearValue})
				frontPtr += 1
				rearPtr -= 1
				// Skip same element
				for frontPtr < rearPtr && input[frontPtr] == input[frontPtr-1] {
					frontPtr += 1
				}
				for frontPtr < rearPtr && input[rearPtr] == input[rearPtr+1] {
					rearPtr -= 1
				}
			} else if baseValue+frontValue+rearValue < 0 {
				frontPtr += 1
			} else {
				rearPtr -= 1
			}
		}
		basePtr += 1
		frontPtr = basePtr + 1
		rearPtr = len(input) - 1
	}
	return output
}
