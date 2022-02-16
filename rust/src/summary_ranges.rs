/*
    https://leetcode.com/problems/summary-ranges/
*/
struct Solution;

impl Solution {
    pub fn summary_ranges(nums: Vec<i32>) -> Vec<String> {
        if nums.is_empty() {
            return vec![];
        }

        let mut result = Vec::new();
        let mut begin = nums[0];
        let mut end = nums[0];

        for &val in nums.iter().skip(1) {
            if end + 1 != val {
                Solution::record(&mut result, begin, end);
                begin = val;
            }
            end = val;
        }

        Solution::record(&mut result, begin, end);

        result
    }

    fn record(vec: &mut Vec<String>, begin: i32, end: i32) {
        if begin == end {
            vec.push(begin.to_string());
        } else {
            vec.push(format!("{}->{}", begin.to_string(), end.to_string()));
        }
    }
}

#[test]
fn test() {
    assert_eq!(
        vec!["0->2", "4->5", "7"],
        Solution::summary_ranges(vec![0, 1, 2, 4, 5, 7])
    );
    assert_eq!(
        vec!["0", "2->4", "6", "8->9"],
        Solution::summary_ranges(vec![0, 2, 3, 4, 6, 8, 9])
    );
}
