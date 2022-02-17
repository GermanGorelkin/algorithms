/*
    https://leetcode.com/problems/triangle/
*/

use std::cmp;
struct Solution;

impl Solution {
    pub fn minimum_total(triangle: Vec<Vec<i32>>) -> i32 {
        let rows = triangle.len();
        let mut dp = vec![vec![]; rows];
        dp[0].push(triangle[0][0]);

        let mut minimum = dp[0][0];

        for row in 1..rows {
            let cols = triangle[row].len();
            for col in 0..cols {
                let (mut cost1, mut cost2) = (i32::MAX, i32::MAX);
                if col < cols - 1 {
                    cost1 = dp[row - 1][col];
                }
                if col > 0 {
                    cost2 = dp[row - 1][col - 1];
                }
                dp[row].push(triangle[row][col] + cmp::min(cost1, cost2));

                if row == rows - 1 {
                    if col == 0 {
                        minimum = dp[row][col]
                    } else {
                        minimum = cmp::min(minimum, dp[row][col])
                    }
                }
            }
        }

        minimum
    }
}

#[test]
fn test() {
    assert_eq!(
        11,
        Solution::minimum_total(vec![vec![2], vec![3, 4], vec![6, 5, 7], vec![4, 1, 8, 3]])
    );
    assert_eq!(-10, Solution::minimum_total(vec![vec![-10]]));
}

/*
Input: triangle = [[2],[3,4],[6,5,7],[4,1,8,3]]
Output: 11
Explanation: The triangle looks like:
   2
  3 4
 6 5 7
4 1 8 3
The minimum path sum from top to bottom is 2 + 3 + 5 + 1 = 11 (underlined above).
*/
