package datastructuresandalgorithms

// To solve this problem, we use 'n' to populate the new array at two different
// positions simultaneously, since the result array has a size of 2*n.
// We use two pointers: 'i' iterates through the first half of the input array (up to n),
// while 'j' acts as a write pointer that advances by two in each iteration
// to interleave the elements correctly without overwriting.

func shuffle(nums []int, n int) []int {
	ra := make([]int, n*2)
	j := 0
	for i := 0; i < n; i++ {
		ra[j] = nums[i]
		ra[j+1] = nums[i+n]
		j = j + 2
	}

	return ra
}

//Another possible solution with a faster runtime

// We pre-allocate a slice with a length of 0 but a capacity of 2n (len(nums)).
// This prevents multiple memory reallocations during the loop.
// We iterate through the first half of the array (up to n).
// In each iteration, we append nums[i] (x) followed by nums[n+i] (y).
// Since append always adds to the end of the slice, the elements are naturally interleaved

func shuffleOptimized(nums []int, n int) []int {
	res := make([]int, 0, n*2)
	for i := 0; i < n; i++ {
		res = append(res, nums[i])
		res = append(res, nums[n+i])
	}
	return res
}
