/*

1.two sum

Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution.

Example:
Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
*/

package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{1, 7, 11, 3, 5, 7, 29, 8}, 30))
}
func twoSum(nums []int, target int) []int {
	mapNums := make(map[int]int)
	for i, val := range nums {
		if j, ok := mapNums[target-val]; ok {
			return []int{j, i}
		}
		mapNums[val] = i
	}
	return []int{0, 0}
}
