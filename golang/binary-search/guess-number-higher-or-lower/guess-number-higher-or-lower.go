/*
We are playing the Guess Game. The game is as follows:

I pick a number from 1 to n. You have to guess which number I picked.

Every time you guess wrong, I'll tell you whether the number is higher or lower.

You call a pre-defined API guess(int num) which returns 3 possible results (-1, 1, or 0):

-1 : My number is lower
 1 : My number is higher
 0 : Congrats! You got it!
Example :

Input: n = 10, pick = 6
Output: 6
*/

package guess_number_higher_or_lower

func guessNumber(n int) int {
	lo, hi := 1, n

	for lo <= hi {
		mid := (lo + hi) / 2
		res := guess(mid)
		if res == 0 {
			return mid
		} else if res == 1 {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}

	return -1
}

var (
	PICK int
)

func guess(num int) int {
	if PICK < num {
		return -1
	}
	if PICK > num {
		return 1
	}
	return 0
}

//type GuessGame struct {
//	pick int
//	n int
//}
//
//func NewGuessGame(pick, n int) (GuessGame,error) {
//	if n < pick {
//		return GuessGame{}, errors.New("error n less pick")
//	}
//
//	return GuessGame{
//		pick: pick,
//		n:    n,
//	}, nil
//}
//
//func (gg GuessGame) guess(num int) int {
//	if gg.pick < num {
//		return -1
//	}
//	if gg.pick > num {
//		return 1
//	}
//
//	return 0
//}
