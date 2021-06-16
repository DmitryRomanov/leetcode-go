package firstbad

import "math"

var (
	badVersion int
)

func firstBadVersion(n int) int {
	left := 0
	right := n
	currentBad := n

	for (right - left) > 1 {
		medianVaule := int(math.Floor(float64((left + right) / 2)))
		isBadVersion := isBadVersion(medianVaule)

		if !isBadVersion {
			left = medianVaule
		} else {
			right = medianVaule
			currentBad = medianVaule
		}
	}
	return currentBad
}

func isBadVersion(version int) bool {
	return badVersion == version
}
