/*
You are given coins of different denominations and a total amount of money amount. Write a function to compute the fewest number of coins that you need to make up that amount. If that amount of money cannot be made up by any combination of the coins, return -1.

Example 1:

Input: coins = [1, 2, 5], amount = 11
Output: 3
Explanation: 11 = 5 + 5 + 1
Example 2:

Input: coins = [2], amount = 3
Output: -1
Note:
You may assume that you have an infinite number of each kind of coin.
*/

package coin_change

/*
1)
Решаем задачу о минимальном числе монет, необходимом представить сумму,
для каждого числа [0, amount].
Получаем ответ для текущего числа на основании ответом для предыдущих чисел.

minCoins[x] - минимальное количество монет из набора coins,
необходимое для получения ровно значения (суммы) x.

minCoins[0] = 0

minCoins[x > 0] = min( minCoins[x - coins[i]] + 1), i = 0..coins.size-1,
вабор только по вариантам x >= coins[i].

Если ни одной монетки нельзя вычислить - присваиваем значение +Inf.

O(amount * coins.size) - time
O(amount) - mem

-------
Example 1:

Input: coins = [1, 2, 5], amount = 11
 0	 1	 2	 3	 4	 5	 6	 7	 8	 9	10	11
[0] [1] [1] [2] [2] [1] [2] [2] [3] [3] [2] [3]

*/
const maxInt = int(^uint(0)>>1) - 1

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0

	for i := 1; i <= amount; i++ {
		dp[i] = maxInt
		for _, coin := range coins {
			if coin <= i {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}

	if dp[amount] == maxInt {
		return -1
	}
	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
