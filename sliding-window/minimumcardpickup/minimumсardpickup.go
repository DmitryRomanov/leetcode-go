// https://leetcode.com/problems/minimum-consecutive-cards-to-pick-up/
package minimumcardpickup

func minimumCardPickup(cards []int) int {
	left := 0
	right := left + 1
	result := len(cards)

	for left < len(cards) {
		if cards[left] == cards[right] {
			if cards[right]-cards[left] < result {
				result = cards[right] - cards[left]
			}
		}
		right++
		if right == len(cards) {
			left++
		}

		if right == left {
			left++
			right = left + 1
		}
	}

	return result
}
