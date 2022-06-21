package peakFinder

import (
	"fmt"
	"testing"
)

type question struct {
	params
	ans
}

type params struct {
	nums []int
}

type ans struct {
	peaks []int
}

func checkIfExists(num int, peaks []int) bool {
	for _, a := range peaks {
		if num == a {
			return true
		}
	}
	return false
}

var qs []question = []question{
	{
		params{[]int{1, 2, 1, 0}},
		ans{[]int{1}},
	},
	{
		params{[]int{1, 2, 3, 4}},
		ans{[]int{3}},
	},
	{
		params{[]int{0, 1}},
		ans{[]int{1}},
	},
	{
		params{[]int{1, 0}},
		ans{[]int{0}},
	},
	{
		params{[]int{1, 2, 1, 4}},
		ans{[]int{1, 3}},
	},
	{
		params{[]int{1, 1, 1, 1}},
		ans{[]int{0, 1, 2, 3}},
	},
	{
		params{[]int{0}},
		ans{[]int{0}},
	},
}

func Test_Solutions(t *testing.T) {
	for _, q := range qs {
		solutionToCheck := StraightForwardSolution(q.nums)
		fmt.Printf("Testing StraightForward implementation for array: %d\n", q.nums)
		if !checkIfExists(solutionToCheck, q.ans.peaks) {
			t.Errorf("Your solution: %d for given array %d, doesn't match any of the possible solutions %d", solutionToCheck, q.nums, q.ans.peaks)
		} else {
			fmt.Printf("Your solution: %d, possible solutions %d\n", solutionToCheck, q.ans.peaks)
		}
	}
	for _, q := range qs {
		solutionToCheck := startFromMiddleSolution(q.nums)
		fmt.Printf("Testing startFromMiddle implementation for array: %d\n", q.nums)
		if !checkIfExists(solutionToCheck, q.ans.peaks) {
			t.Errorf("Your solution: %d for given array %d, doesn't match any of the possible solutions %d", solutionToCheck, q.nums, q.ans.peaks)
		} else {
			fmt.Printf("Your solution: %d, possible solutions %d\n", solutionToCheck, q.ans.peaks)
		}
	}
}
