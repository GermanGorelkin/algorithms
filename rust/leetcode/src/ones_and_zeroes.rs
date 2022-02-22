/*
    https://leetcode.com/problems/ones-and-zeroes/
*/
struct Solution;

use std::cmp;

impl Solution {
    pub fn find_max_form(strs: Vec<String>, m: i32, n: i32) -> i32 {
        let (m, n) = (m as usize, n as usize);
        let mut dp = vec![vec![0; n + 1]; m + 1];

        for s in &strs {
            let (mut ones, mut zeros) = (0, 0);
            s.chars().for_each(|ch| {
                if ch == '1' {
                    ones += 1;
                } else {
                    zeros += 1;
                }
            });

            for i in (zeros..=m).rev() {
                for j in (ones..=n).rev() {
                    dp[i][j] = cmp::max(dp[i][j], dp[i - zeros][j - ones] + 1);
                }
            }
        }

        dp[m][n]
    }
}

#[test]
fn test() {
    assert_eq!(
        4,
        Solution::find_max_form(
            vec![
                "10".to_string(),
                "0001".to_string(),
                "111001".to_string(),
                "1".to_string(),
                "0".to_string()
            ],
            5,
            3
        )
    );
    assert_eq!(
        2,
        Solution::find_max_form(
            vec!["10".to_string(), "0".to_string(), "1".to_string()],
            1,
            1
        )
    );
}
