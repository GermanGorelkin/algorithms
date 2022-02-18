use std::cmp;

/*
    https://leetcode.com/problems/maximal-square/
*/
struct Solution;

impl Solution {
    pub fn maximal_square(matrix: Vec<Vec<char>>) -> i32 {
        let rows = matrix.len();
        let cols = matrix[0].len();
        let mut dp = vec![vec![0; cols + 1]; rows + 1];
        let mut max_len: i32 = 0;

        for row in 1..=rows {
            for col in 1..=cols {
                if matrix[row - 1][col - 1] == '1' {
                    dp[row][col] = 1 + cmp::min(
                        cmp::min(dp[row - 1][col - 1], dp[row - 1][col]),
                        dp[row][col - 1],
                    );
                    max_len = cmp::max(max_len, dp[row][col] as i32);
                }
            }
        }

        max_len * max_len
    }
}

#[test]
fn test() {
    assert_eq!(
        4,
        Solution::maximal_square(vec![
            vec!['1', '0', '1', '0', '0'],
            vec!['1', '0', '1', '1', '1'],
            vec!['1', '1', '1', '1', '1'],
            vec!['1', '0', '0', '1', '0'],
        ])
    );
    assert_eq!(
        1,
        Solution::maximal_square(vec![vec!['0', '1'], vec!['1', '0']])
    );
    assert_eq!(1, Solution::maximal_square(vec![vec!['1']]));
    assert_eq!(0, Solution::maximal_square(vec![vec!['0']]));
}
