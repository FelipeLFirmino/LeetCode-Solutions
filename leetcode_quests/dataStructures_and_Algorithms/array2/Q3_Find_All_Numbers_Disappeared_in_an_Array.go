package datastructuresandalgorithms

// Time Complexity: O(n)
// Space Complexity: O(n) - Because of the 'freq' array allocation.

// This approach uses an auxiliary array 'freq' to keep track of seen numbers.
// We iterate through the input and increment the count at freq[value-1].
// We then iterate through the 'freq' array; any index 'i' with a value of 0
// means the number 'i+1' was missing from the original input.

func findDisappearedNumbers(nums []int) []int {
	k := len(nums)
	freq := make([]int, k)
	// Count frequencies:
	for _, x := range nums {
		freq[x-1]++
	}
	// Loop to know which number has freq 0
	res := make([]int, 0)
	for i := 0; i < k; i++ {
		if freq[i] == 0 {
			res = append(res, i+1)
		}
	}
	return res
}

// Time Complexity: O(n)
// Space Complexity: O(1) - Modifies the input array to save space.

// This strategy uses the sign bit of the numbers already in the array as a "flag".
// For every number 'x' in 'nums', we go to the index 'abs(x)-1'.
// We flip the value at that index to negative: nums[abs(x)-1] = -abs(nums[abs(x)-1]).
// This effectively marks that the number 'index+1' has been encountered.
// We use the Absolute Value (Abs) because a number might have been flipped
// to negative by a previous step, but we still need its original value to find the next index.
// In the final pass, any index 'i' that still contains a POSITIVE number
// was never "visited", meaning the number 'i+1' is missing.
func IntegerAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func findDisappearedNumbersOptimized(nums []int) []int {
	for _, number := range nums {
		presentIndex := IntegerAbs(number) - 1
		if nums[presentIndex] > 0 {
			// use negative number to mark number is presented in array
			nums[presentIndex] = -nums[presentIndex]
		}
	}

	// the disappeared numbers are those grids which are still with positive value
	missingNumber := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			missingNumber = append(missingNumber, i+1)
		}
	}

	return missingNumber
}
