package main

import "fmt"

/*

Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.
You can return the answer in any order.

Example 1:

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

Example 2:

Input: nums = [3,2,4], target = 6
Output: [1,2]
Example 3:

Input: nums = [3,3], target = 6
Output: [0,1]


Constraints:

2 <= nums.length <= 10^4
-10^9 <= nums[i] <= 10^9
-10^9 <= target <= 10^9
Only one valid answer exists.


Follow-up: Can you come up with an algorithm that is less than O(n2) time complexity?

*/

// twoSumBruteForceSolution - this is a brute force solution
// Time: O(n^2)
// Space: O(1)
func twoSumBruteForceSolution(nums []int, target int) []int {

	// compare each element with next element in the nums array
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			// pair found
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	// no pair found
	return []int{-1, -1}
}

func main() {

	// O(n^2) test cases
	// case 1
	fmt.Println("Test case 1: []nums=[2,7,11,15], target=9, Expected Output: []result=[0,1] => Output: ", twoSumBruteForceSolution([]int{2, 7, 11, 15}, 9))

	// case 2
	fmt.Println("Test case 2: []nums=[3,2,4], target=6, Expected Output: []result=[1,2] => Output: ", twoSumBruteForceSolution([]int{3, 2, 4}, 6))

	// case 3
	fmt.Println("Test case 3: []nums=[3,3], target=6, Expected Output: []result=[0,1] => Output: ", twoSumBruteForceSolution([]int{3, 3}, 6))

	// case 4
	fmt.Println("Test case 4: []nums=[3,3,3], target=7, Expected Output: []result=[-1,-1] => Output: ", twoSumBruteForceSolution([]int{3, 3, 3}, 7))

}
