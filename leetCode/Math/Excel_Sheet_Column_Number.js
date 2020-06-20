/**
Given a column title as appear in an Excel sheet, return its corresponding column number.

For example:

    A -> 1
    B -> 2
    C -> 3
    ...
    Z -> 26
    AA -> 27
    AB -> 28
    ...
Example 1:

Input: "A"
Output: 1
Example 2:

Input: "AB"
Output: 28
Example 3:

Input: "ZY"
Output: 701
 */

/**
 * @param {string} s
 * @return {number}
 */
const a = '0ABCDEFGHIJKLMNOPQRSTUVWXYZ'

const titleToNumber = (s) => {
  return Array
    .from(s)
    .reverse()
    .reduce((sum, v, i) => sum + a.indexOf(v) * Math.pow(26, i), 0)
}

/**
 * time complexity O(Length of s)
 */
