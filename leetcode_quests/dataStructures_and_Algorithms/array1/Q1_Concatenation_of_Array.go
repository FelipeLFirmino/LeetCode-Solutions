package datastructuresandalgorithms

//in this problem we are able to fill both right spots at the same time by adding the lenght of n to the index
//since the final array is always twice its size
func getConcatenation(nums []int) []int {
	n := len(nums)
	ans := make([]int, 2*n)

	for i := 0; i < n; i++ {
		ans[i] = nums[i]
		ans[i+n] = nums[i]
	}
	return ans
}
