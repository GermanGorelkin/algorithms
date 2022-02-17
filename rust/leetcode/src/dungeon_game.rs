use std::cmp;

/*
    https://leetcode.com/problems/dungeon-game/
*/
struct Solution;

impl Solution {
    pub fn calculate_minimum_hp(dungeon: Vec<Vec<i32>>) -> i32 {
        let rows = dungeon.len();
        let cols = dungeon[0].len();

        let mut dp = vec![vec![i32::MAX; cols + 1]; rows + 1];
        dp[rows - 1][cols] = 1;
        dp[rows][cols - 1] = 1;

        for row in (0..rows).rev() {
            for col in (0..cols).rev() {
                let val = cmp::min(dp[row + 1][col], dp[row][col + 1]);
                dp[row][col] = cmp::max(1, val - dungeon[row][col]);
            }
        }

        dp[0][0]
    }
}

/*
-2, -3, 3
-5, -10, 1
10, 30, -5

7, 5,  2, m
6, 15, 5, m
1, 1,  6, 1
m, m, 1, m
*/

#[test]
fn test() {
    assert_eq!(
        7,
        Solution::calculate_minimum_hp(vec![vec![-2, -3, 3], vec![-5, -10, 1], vec![10, 30, -5]])
    );
    assert_eq!(1, Solution::calculate_minimum_hp(vec![vec![0]]));
}
