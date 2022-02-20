use std::cmp;

/*
    https://leetcode.com/problems/perfect-squares/
*/
struct Solution;

impl Solution {
    pub fn num_squares(n: i32) -> i32 {
        let mut dp = vec![i32::MAX; (n + 1) as usize];
        dp[0] = 0;

        for i in (1..).take_while(|x| x * x <= n) {
            let sq = (i * i) as usize;
            for j in sq..=n as usize {
                dp[j] = cmp::min(dp[j], dp[j - sq] + 1);
            }
        }

        dp[n as usize]
    }
}

/*
    |0 1 2 3 4 5 6 7 8 9 10 11 12
   1|  1 2 3 4 5 6 7 8 9 10 11 12
 2*2|        1 2 3 4 2 3 4  5  3
 3*3|                  1 2  3  4
*/

#[test]
fn test() {
    assert_eq!(3, Solution::num_squares(12));
    assert_eq!(2, Solution::num_squares(13));
}
