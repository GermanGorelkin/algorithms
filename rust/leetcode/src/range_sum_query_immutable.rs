/*
    https://leetcode.com/problems/range-sum-query-immutable/
*/

struct NumArray {
    nums: Vec<i32>,
}

impl NumArray {
    fn new(nums: Vec<i32>) -> Self {
        NumArray { nums }
    }

    fn sum_range(&self, left: i32, right: i32) -> i32 {
        let mut sum: i32 = 0;

        self.nums[left as usize..=right as usize]
            .iter()
            .for_each(|&n| {
                sum += n;
            });

        sum
    }
}

#[test]
fn test() {
    let nums = vec![-2, 0, 3, -5, 2, -1];
    let obj = NumArray::new(nums);

    assert_eq!(1, obj.sum_range(0, 2));
    assert_eq!(-1, obj.sum_range(2, 5));
    assert_eq!(-3, obj.sum_range(0, 5));
}
