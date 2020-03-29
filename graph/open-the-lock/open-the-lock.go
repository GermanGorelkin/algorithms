/*
You have a lock in front of you with 4 circular wheels. Each wheel has 10 slots: '0', '1', '2', '3', '4', '5', '6', '7', '8', '9'. The wheels can rotate freely and wrap around: for example we can turn '9' to be '0', or '0' to be '9'. Each move consists of turning one wheel one slot.

The lock initially starts at '0000', a string representing the state of the 4 wheels.

You are given a list of deadends dead ends, meaning if the lock displays any of these codes, the wheels of the lock will stop turning and you will be unable to open it.

Given a target representing the value of the wheels that will unlock the lock, return the minimum total number of turns required to open the lock, or -1 if it is impossible.

Example 1:
Input: deadends = ["0201","0101","0102","1212","2002"], target = "0202"
Output: 6
Explanation:
A sequence of valid moves would be "0000" -> "1000" -> "1100" -> "1200" -> "1201" -> "1202" -> "0202".
Note that a sequence like "0000" -> "0001" -> "0002" -> "0102" -> "0202" would be invalid,
because the wheels of the lock become stuck after the display becomes the dead end "0102".
Example 2:
Input: deadends = ["8888"], target = "0009"
Output: 1
Explanation:
We can turn the last wheel in reverse to move from "0000" -> "0009".
Example 3:
Input: deadends = ["8887","8889","8878","8898","8788","8988","7888","9888"], target = "8888"
Output: -1
Explanation:
We can't reach the target without getting stuck.
Example 4:
Input: deadends = ["0000"], target = "8888"
Output: -1
Note:
The length of deadends will be in the range [1, 500].
target will not be in the list deadends.
Every string in deadends and the string target will be a string of 4 digits from the 10,000 possibilities '0000' to '9999'.
*/

package open_the_lock

func openLock(deadends []string, target string) int {
	deadEnds := make(map[string]struct{}, len(deadends))
	for _, de := range deadends {
		deadEnds[de] = struct{}{}
	}
	if _, ok := deadEnds["0000"]; ok {
		return -1
	}

	var visited = make(map[string]struct{})
	var step int
	wheels := make([]string, 8)

	start := "0000"
	q := &queue{}
	q.enQueue(start)
	visited[start] = struct{}{}

	for !q.isEmpty() {
		step++

		for i := q.size(); i > 0; i-- {
			d := q.deQueue()
			genWheels(d, wheels)

			for _, w := range wheels {
				if _, ok := visited[w]; ok {
					continue
				}
				if w == target {
					return step
				}

				if _, ok := deadEnds[w]; ok {
					continue
				}

				visited[w] = struct{}{}
				q.enQueue(w)
			}
		}
	}

	return -1
}

type queue struct {
	data []string
}

func (q *queue) enQueue(v string) {
	q.data = append(q.data, v)
}
func (q *queue) deQueue() string {
	v := q.data[0]
	q.data = q.data[1:]
	return v
}
func (q *queue) isEmpty() bool {
	return len(q.data) == 0
}
func (q *queue) size() int {
	return len(q.data)
}

func genWheels(state string, wheels []string) {
	wheels = wheels[:0]
	for i := range state {
		d1, d2 := incCh(state[i])

		b := []byte(state)
		b[i] = d1
		wheels = append(wheels, string(b))

		b = []byte(state)
		b[i] = d2
		wheels = append(wheels, string(b))
	}
}

func incCh(ch byte) (byte, byte) {
	var d1, d2 byte

	if ch == '0' {
		d1 = '9'
	} else {
		d1 = ch - 1
	}

	if ch == '9' {
		d2 = '0'
	} else {
		d2 = ch + 1
	}

	return d1, d2
}
