// https://leetcode.com/problems/minimum-consecutive-cards-to-pick-up/
package minimumcardpickup

func minimumCardPickup(cards []int) int {
	result := len(cards)
	found := false

	card_positions := make(map[int]int, len(cards))
	for i := 0; i < len(cards); i++ {
		card_value := cards[i]
		if last_position, exists := card_positions[card_value]; exists {
			current_result := i - last_position + 1
			if current_result < result {
				result = current_result
			}
			found = true
		}
		card_positions[card_value] = i
	}

	if found {
		return result
	} else {
		return -1
	}
}

func minimumCardPickupBruteforce(cards []int) int {
	result := len(cards)
	found := false

	for left := 0; left < len(cards); left++ {
		for right := left + 1; right < len(cards); right++ {
			if cards[left] == cards[right] {
				found = true
				curResult := right - left + 1
				if right-left < result {
					result = curResult
				}
			}
		}
	}

	if found {
		return result
	} else {
		return -1
	}
}
