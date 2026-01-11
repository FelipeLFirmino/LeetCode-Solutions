package datastructuresandalgorithms

// To solve this, we must consider things:
// 1. The values in 'nums' are never 0.
// 2. The array is NOT necessarily sorted, even if examples suggest otherwise.

// In this solution, we iterate through 'nums' and use its VALUES as INDEXES
// to populate a 'seen' array.
// If we encounter an index that already has a non-zero (!= 0) value, we've found the duplicate.
// The duplicate entry causes one specific index in the 'seen' array to remain 0,
// which represents the missing number.
// Finally, we iterate through 'seen' to find that 0 and return its index + 1.
func findErrorNums(nums []int) []int {

	seen := make([]int, len(nums))
	var res []int

	for _, n := range nums {
		if seen[n-1] == 0 {
			seen[n-1] = n
		} else {
			res = append(res, n)
		}
	}

	for i, n := range seen {
		if n == 0 {
			res = append(res, i+1)
		}
	}

	return res

}
