package datastructuresandalgorithms

// This solution has O(n^2) complexity, but is the simplest.
// By using a nested loop, the execution time grows quadratically as the array size increases.
// Although it is the simplest approach, comparing each element against the entire array
// is inefficient for large datasets.

func smallerNumbersThanCurrent(nums []int) []int {
	output := make([]int, 0, len(nums))

	for _, e := range nums {
		count := 0
		for _, i := range nums {
			if e > i {
				count++
			}
		}
		output = append(output, count)
	}
	return output
}

// This solution has O(n) time complexity.
// The constraint 0 <= nums[i] <= 100 allows us to use a frequency array of constant size (K=101).

// First loop: Count how many times each number appears (Frequency array).

// Second loop: Calculate the prefix sum (prefix[v] = prefix[v-1] + freq[v-1]). Each index 'v' in the prefix array
// now stores the total count of numbers smaller than 'v'.

// Third loop: Map these pre-calculated counts back to the original positions
// of the elements in 'nums' to produce the final output.

func smallerNumbersThanCurrentOptmized(nums []int) []int {
	const K = 101
	freq := make([]int, K)
	// Count frequencies:
	for _, x := range nums {
		freq[x]++
	}
	// Build prefix sum:
	prefix := make([]int, K)
	for v := 1; v < K; v++ {
		prefix[v] = prefix[v-1] + freq[v-1]
	}
	// Answer for each num:
	res := make([]int, len(nums))
	for i, x := range nums {
		res[i] = prefix[x]
	}
	return res
}
