/*
    https://leetcode.com/problems/missing-number/
*/
struct Solution;

impl Solution {
    pub fn missing_number(nums: Vec<i32>) -> i32 {
        let mut result = nums.len() as i32;
        for (i, &val) in nums.iter().enumerate() {
            result = result ^ i as i32 ^ val; // a^b^b=a
        }
        return result;
    }
}

#[test]
fn test() {
    assert_eq!(2, Solution::missing_number(vec![3, 0, 1]));
    assert_eq!(2, Solution::missing_number(vec![0, 1]));
    assert_eq!(8, Solution::missing_number(vec![9, 6, 4, 2, 3, 5, 7, 0, 1]));
}
