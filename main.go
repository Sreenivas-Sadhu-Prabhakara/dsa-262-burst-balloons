package main

import "fmt"

// MaxCoins returns the most coins from bursting all balloons, where bursting
// balloon i pays left·i·right for its CURRENT neighbors. Interval DP keyed
// on the LAST balloon burst in each range, with virtual 1-balloons padding
// the ends.
func MaxCoins(nums []int) int {
	n := len(nums)
	a := make([]int, n+2)
	a[0], a[n+1] = 1, 1
	copy(a[1:], nums)
	// dp[l][r]: best for bursting everything strictly between l and r.
	dp := make([][]int, n+2)
	for i := range dp {
		dp[i] = make([]int, n+2)
	}
	for width := 2; width <= n+1; width++ {
		for l := 0; l+width <= n+1; l++ {
			r := l + width
			for k := l + 1; k < r; k++ { // k bursts LAST: its neighbors are l and r
				coins := dp[l][k] + dp[k][r] + a[l]*a[k]*a[r]
				if coins > dp[l][r] {
					dp[l][r] = coins
				}
			}
		}
	}
	return dp[0][n+1]
}

func main() {
	fmt.Println(MaxCoins([]int{3, 1, 5, 8}))
}
