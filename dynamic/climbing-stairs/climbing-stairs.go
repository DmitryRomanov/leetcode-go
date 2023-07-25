// https://leetcode.com/problems/climbing-stairs/description/
package climbing_stairs

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	n1, n2 := 0, 1
	for i := 0; i < n; i++ {
		n1, n2 = n2, n1+n2
	}
	return n2
}
