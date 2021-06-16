package firstbad

import "math"

var (
	badVersion int
)

func firstBadVersion(n int) int {
	left := 0
	right := n

	for (right - left) > 0 {
		medianVaule := int(math.Floor(float64((left + right) / 2)))
		isBadVersion := isBadVersion(medianVaule)

		if isBadVersion {
			return medianVaule
		}
	}
	return -1
}

func isBadVersion(version int) bool {
	return badVersion == version
}
