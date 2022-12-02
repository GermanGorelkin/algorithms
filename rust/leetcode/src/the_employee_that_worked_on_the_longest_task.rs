/*
    https://leetcode.com/problems/the-employee-that-worked-on-the-longest-task/
*/

struct Solution;

impl Solution {
    pub fn hardest_worker(n: i32, logs: Vec<Vec<i32>>) -> i32 {
        let mut result = (logs[0][1], logs[0][0]); // (longest time, smallest emplId)

        for i in 1..logs.len() {
            let time = logs[i][1] - logs[i - 1][1];

            if time > result.0 {
                result = (time, logs[i][0]);
            } else if time == result.0 {
                let emplId = logs[i][0];
                if result.1 > emplId {
                    result.1 = emplId;
                }
            }
        }

        result.1
    }
}

#[test]
fn test() {
    assert_eq!(
        1,
        Solution::hardest_worker(10, vec![vec![0, 3], vec![2, 5], vec![0, 9], vec![1, 15]])
    );
    assert_eq!(
        3,
        Solution::hardest_worker(26, vec![vec![1, 1], vec![3, 7], vec![2, 12], vec![7, 17]])
    );
    assert_eq!(
        0,
        Solution::hardest_worker(2, vec![vec![0, 10], vec![1, 20]])
    );
}
