use std::cmp;

/*
    https://leetcode.com/problems/minimum-path-sum/
*/
struct Solution;

impl Solution {
    pub fn min_path_sum(grid: Vec<Vec<i32>>) -> i32 {
        let rows = grid.len();
        let cols = grid.get(0).unwrap().len();
        let mut dp: Vec<Vec<i32>> = vec![vec![0; cols]; rows];

        for i in 0..rows {
            for j in 0..cols {
                let (mut cost1, mut cost2) = (i32::MAX, i32::MAX);
                if i > 0 {
                    cost1 = dp[i - 1][j];
                }
                if j > 0 {
                    cost2 = dp[i][j - 1];
                }
                dp[i][j] = grid[i][j];
                if cost1 != i32::MAX || cost2 != i32::MAX {
                    dp[i][j] += cmp::min(cost1, cost2);
                }
            }
        }
        dp[rows - 1][cols - 1]
    }
}

#[test]
fn test() {
    assert_eq!(
        7,
        Solution::min_path_sum(vec![vec![1, 3, 1], vec![1, 5, 1], vec![4, 2, 1]])
    );
    assert_eq!(
        12,
        Solution::min_path_sum(vec![vec![1, 2, 3], vec![4, 5, 6]])
    );
}
