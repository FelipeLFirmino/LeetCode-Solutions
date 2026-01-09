package main

//To solve remove element we need to think of k as a counter and index at the same time
//K IS THE NUMBER OF VALUES != VAL
//Once we find numbers that are differ from val we overwrite them, and add one to k
//so the index moves as numbers diffent from k are being found, THE NUMBERS  AT AND AFTER nums[k] does not matter
func removeElement(nums []int, val int) int {
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}
