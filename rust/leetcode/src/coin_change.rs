use std::cmp;

/*
    https://leetcode.com/problems/coin-change/
*/
struct Solution;

impl Solution {
    pub fn coin_change(coins: Vec<i32>, amount: i32) -> i32 {
        let mut dp = vec![amount + 1; (amount + 1) as usize];
        dp[0] = 0;

        for i in 1..=amount {
            for &coin in coins.iter().filter(|&&c| c <= i) {
                dp[i as usize] = cmp::min(dp[i as usize], dp[(i - coin) as usize] + 1);
            }
        }

        if dp[amount as usize] > amount {
            -1
        } else {
            dp[amount as usize]
        }
    }
}

/*
    [1, 2, 5], 11

    [0, 1, 1, 1, 2, 1, 2, 2, 2, 3, 2, 3]
*/

#[test]
fn test() {
    assert_eq!(3, Solution::coin_change(vec![1, 2, 5], 11));
    assert_eq!(20, Solution::coin_change(vec![186, 419, 83, 408], 6249));
    assert_eq!(-1, Solution::coin_change(vec![2], 3));
    assert_eq!(0, Solution::coin_change(vec![1], 0));
}
