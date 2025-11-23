package leetcode

// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
// You can return the answer in any order.

// Example 1:
// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
// Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

// Example 2:
// Input: nums = [3,2,4], target = 6
// Output: [1,2]

// Example 3:
// Input: nums = [3,3], target = 6
// Output: [0,1]

// Constraints:
// 2 <= nums.length <= 104
// -109 <= nums[i] <= 109
// -109 <= target <= 109
// Only one valid answer exists.

func TwoSum(nums []int, target int) map[string]func() any {
	fns := map[string]func() any{
		"Hashmap (The Best)":      func() any { return twoSumHashmap(nums, target) },
		"Brute Force (The Worst)": func() any { return twoSumBruteForce(nums, target) },
	}

	return fns
}

func twoSumHashmap(nums []int, target int) []int {
	hashmap := make(map[int]int, len(nums))
	for i, val := range nums {
		if pos, ok := hashmap[target-val]; ok {
			return []int{pos, i}
		}
		hashmap[val] = i
	}

	return []int{}
}

func twoSumBruteForce(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{}
}
