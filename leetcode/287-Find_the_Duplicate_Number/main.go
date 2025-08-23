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
