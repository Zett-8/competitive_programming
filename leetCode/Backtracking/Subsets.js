/**
Given a set of distinct integers, nums, return all possible subsets (the power set).

Note: The solution set must not contain duplicate subsets.

Example:

Input: nums = [1,2,3]
Output:
[
  [3],
  [1],
  [2],
  [1,2,3],
  [1,3],
  [2,3],
  [1,2],
  []
]
 */

/**
 * @param {number[]} nums
 * @return {number[][]}
 */

const subsets = (nums) => {
  const ans = []

  const r = (arr, index) => {
    ans.push(arr)
    for (let i = index; i<nums.length; i++) {
      r([...arr, nums[i]], i+1)
    }
  }
  r([], 0)

  return ans
}

/**
 * time complexity O(2^n)
 */
