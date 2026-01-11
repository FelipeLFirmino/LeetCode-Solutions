package datastructuresandalgorithms

//This problema is simple, we just need to variables to count the
// Current Count(c) and the  max Count number of ones found (mc)
//once we find a 0 we restart the counter since the next one found will not be consecutive
// if we find a 1 we also have to check if c became bigger than mc so we can store our new max count

func findMaxConsecutiveOnes(nums []int) int {
	c, mc := 0, 0

	for _, v := range nums {
		if v == 1 {
			c++
			if c > mc {
				mc = c
			}
		} else {
			c = 0
		}
	}
	return mc
}
