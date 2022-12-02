struct Solution;

impl Solution {
    pub fn check_perfect_number(num: i32) -> bool {
        if num <= 1 {
            return false;
        }
        let mut result = 0;

        for n in 1..=(num as f64).sqrt() as i32 {
            if num % n == 0 {
                result += n;

                if n != 1 && n * n != num {
                    result += num / n;
                }
            }
        }

        return result == num;
    }
}

#[test]
fn test() {
    assert_eq!(true, Solution::check_perfect_number(28));
    assert_eq!(false, Solution::check_perfect_number(7));
    assert_eq!(false, Solution::check_perfect_number(1));
}
