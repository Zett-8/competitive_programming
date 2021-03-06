/**
 You are given coins of different denominations and a total amount of money amount. Write a function to compute the fewest number of coins that you need to make up that amount. If that amount of money cannot be made up by any combination of the coins, return -1.

Example 1:

Input: coins = [1, 2, 5], amount = 11
Output: 3
Explanation: 11 = 5 + 5 + 1
 */

/**
 * @param {number[]} coins
 * @param {number} amount
 * @return {number}
 */

const coinChange = function(coins, amount) {
  coins = coins.sort((a,b) => a - b)
  const dp = new Array(amount+1).fill(amount+1)
  dp[0] = 0

  for (let i = 1; i <= amount; i++) {
    for (let coin of coins) {
      if (i < coin) break
      dp[i] = Math.min(dp[i], dp[i-coin]+1)
    }
  }

  return dp[amount] > amount ? -1 : dp[amount]
}
