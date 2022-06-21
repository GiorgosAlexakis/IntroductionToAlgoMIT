package peakFinder

/*
Problem: Find a peak
0 1 2 3 4 5 6 7 8
a b c d e f g h i
Position 0 is a peak if and only if a >= b
Position 2 is a peak if and only if b >= a and b >= c
Position 8 is a peak if and only if i >= h
*/

func StraightForwardSolution(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if len(nums) == 1 {
			return 0
		}
		if i == 0 && nums[0] >= nums[1] {
			return i
		}
		if i == len(nums)-1 && nums[i] >= nums[i-1] {
			return i
		}
		if i > 0 && i < len(nums)-1 && nums[i] >= nums[i-1] && nums[i] >= nums[i+1] {
			return i
		}
	}
	return 0
}

func startFromMiddleSolution(nums []int) int {
	index := 0
	var b func(int, int)
	b = func(left int, right int) {
		mid := (right-left)/2 + left
		if nums[mid] < nums[mid-1] {
			right = mid - 1
			b(left, right)
		} else if nums[mid+1] > nums[mid] {
			left = mid + 1
			b(left, right)
		} else {
			index = mid
			return
		}
	}
	b(0, len(nums)-1)
	return index
}
