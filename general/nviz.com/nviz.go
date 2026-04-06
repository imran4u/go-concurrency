package main

//Given an array of intervals where intervals[i] = [starti, endi], merge all overlapping intervals, and return an array of the non-overlapping intervals that cover all the intervals in the input.
//Example 1:
// Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
// Output: [[1,6],[8,10],[15,18]]
// Explanation: Since intervals [1,3] and [2,6] overlap, merge them into [1,6].
// Example 2:
// Input: intervals = [[1,4],[4,5]]
// Output: [[1,5]]
// Explanation: Intervals [1,4] and [4,5] are considered overlapping.
// Example 3:
// Input: intervals = [[4,7],[1,4]]
// Output: [[1,7]]
// Explanation: Intervals [1,4] and [4,7] are considered overlapping.

// [[4,7],[8,9], [1,4]]

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	// input := [][]int{{1, 3}, {8, 10}, {2, 6}, {15, 18}}
	// 1,5],[2,3][3,6]

	input := [][]int{{1, 5}, {2, 3}, {3, 6}}

	// Sort by age using slices.SortFunc
	slices.SortFunc(input, func(a, b []int) int {
		return cmp.Compare(a[0], b[0])
	})
	fmt.Println("sorted inut", input)

	result := [][]int{}
	n := len(input)
	p1 := 0
	p2 := 1

	for p2 < n {
		f := input[p1][0]
		if input[p2-1][1] >= input[p2][0] {
			p2++
			continue
		}
		s := input[p2-1][1]
		in := []int{f, s}
		result = append(result, in)
		p1 = p2
		p2++

	}

	// if last element do not have overlap
	if p1 == n-1 {
		result = append(result, input[p1])
	}

	// if there is only one element, or overlap till end
	if p1 == 0 {
		in := []int{input[p1][0], input[p2-1][1]}
		result = append(result, in)
	}

	fmt.Println(result)
}
