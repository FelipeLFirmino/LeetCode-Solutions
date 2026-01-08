package leetcodesolutions

// To solve this problem, we iterate through the arrays from back to front using two indices (one for each array)
// // and a third index to track the current position in the total length of nums1.
// // Every iteration must add a value to the current index of nums1 (starting from the end)
// // always choosing the GREATEST VALUE AVAILABLE in the loop (either from nums1 or nums2).
func merge(nums1 []int, m int, nums2 []int, n int) {
	p1 := m - 1
	p2 := n - 1
	p := m + n - 1
	for p2 >= 0 {
		if p1 >= 0 && nums1[p1] > nums2[p2] {
			nums1[p] = nums1[p1]
			p1--
		} else {
			nums1[p] = nums2[p2]
			p2--
		}
		p--
	}

}
