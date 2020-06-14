/**
 You are given two integer arrays nums1 and nums2 sorted in ascending order and an integer k.

Define a pair (u,v) which consists of one element from the first array and one element from the second array.

Find the k pairs (u1,v1),(u2,v2) ...(uk,vk) with the smallest sums.

Example 1:

Input: nums1 = [1,7,11], nums2 = [2,4,6], k = 3
Output: [[1,2],[1,4],[1,6]]
Explanation: The first 3 pairs are returned from the sequence:
[1,2],[1,4],[1,6],[7,2],[7,4],[11,2],[7,6],[11,4],[11,6]
 */

/**
 * @param {number[]} nums1
 * @param {number[]} nums2
 * @param {number} k
 * @return {number[][]}
 */

const kSmallestPairs = function(nums1, nums2, k) {
  const ans = []

  for (let x = 0; x < nums1.length; x++) {
    for (let y = 0; y < nums2.length; y++) {
      ans.push([nums1[x],nums2[y]])
    }
  }

  return ans.sort((a, b) => (a[0] + a[1]) - (b[0] + b[1])).slice(0, k)
}
