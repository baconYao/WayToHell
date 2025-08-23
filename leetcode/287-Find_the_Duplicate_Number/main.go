package findtheduplicatenumber

// // Solution 1
// func findDuplicate(nums []int) int {
//     freq := make([]int, len(nums))

//     for _, v := range nums {
//         if freq[v-1] == 1 {
//         return v
//         } else {
//         freq[v-1]++
//         }
//     }

//     return 0
// }

func findDuplicate(nums []int) int {
	slow := nums[0]
	fast := nums[0]
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}
	slow = nums[0]
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}

	return slow

}
