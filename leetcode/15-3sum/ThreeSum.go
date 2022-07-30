package threesum

import "sort"

// Method 1
// Runtime: 49 ms
// Memory Usage: 7.4 MB
func ThreeSum(nums []int) [][]int {
	output := make([][]int, 0)
	sort.Ints(nums) // Sort accending
	if len(nums) < 3 {
		return output
	}

	for basePtr := 0; basePtr < len(nums)-2; basePtr++ {
		// Skip duplicate
		if basePtr > 0 && nums[basePtr] == nums[basePtr-1] {
			continue
		}
		frontPtr, rearPtr := basePtr+1, len(nums)-1
		for frontPtr < rearPtr {
			sum := nums[basePtr] + nums[frontPtr] + nums[rearPtr]
			if sum > 0 {
				rearPtr -= 1
			} else if sum < 0 {
				frontPtr += 1
			} else {
				output = append(output, []int{nums[basePtr], nums[frontPtr], nums[rearPtr]})
				frontPtr += 1
				rearPtr -= 1
				// Skip duplicate
				for frontPtr < rearPtr && nums[frontPtr] == nums[frontPtr-1] {
					frontPtr += 1
				}
				for frontPtr < rearPtr && nums[rearPtr] == nums[rearPtr+1] {
					rearPtr -= 1
				}
			}
		}
	}

	return output
}
