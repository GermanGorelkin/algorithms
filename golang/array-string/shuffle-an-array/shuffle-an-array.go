/*
Shuffle a set of numbers without duplicates.

Example:

// Init an array with set 1, 2, and 3.
int[] nums = {1,2,3};
Solution solution = new Solution(nums);

// Shuffle the array [1,2,3] and return its result. Any permutation of [1,2,3] must equally likely to be returned.
solution.shuffle();

// Resets the array back to its original configuration [1,2,3].
solution.reset();

// Returns the random shuffling of array [1,2,3].
solution.shuffle();
*/

package shuffle_an_array

import "math/rand"

type Solution struct {
	array []int
}

func Constructor(nums []int) Solution {
	return Solution{array: nums}
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.array
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	res := append([]int{}, this.array...)
	for i := len(res) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		res[i], res[j] = res[j], res[i]
	}
	return res
}
