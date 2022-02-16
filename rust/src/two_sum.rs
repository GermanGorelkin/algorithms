/*
    https://leetcode.com/problems/two-sum/
*/
struct Solution;

use std::collections::HashMap;

impl Solution {
    pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
        let mut m: HashMap<i32, usize> = HashMap::with_capacity(nums.len());

        for (i, &val) in nums.iter().enumerate() {
            if let Some(&j) = m.get(&(target - val)) {
                return vec![j as i32, i as i32];
            }
            m.insert(val, i);
        }
        return vec![];
    }
}

#[test]
fn test() {
    assert_eq!(vec![0, 1], Solution::two_sum(vec![2, 7, 11, 15], 9));
    assert_eq!(vec![1, 2], Solution::two_sum(vec![3, 2, 4], 6));
}
