package mergesortedarray

func merge(nums1 []int, m int, nums2 []int, n int) {
	p1, p2, p3 := m-1, n-1, m+n-1
	for p1 != -1 || p2 != -1 {
		if p1 >= 0 && p2 >= 0 {
			if nums1[p1] > nums2[p2] {
				nums1[p3] = nums1[p1]
				p1 -= 1
			} else {
				nums1[p3] = nums2[p2]
				p2 -= 1
			}
			p3 -= 1
		} else if p1 < 0 && p2 >= 0 {
			nums1[p3] = nums2[p2]
			p2 -= 1
			p3 -= 1
		} else {
			break
		}
	}
}
