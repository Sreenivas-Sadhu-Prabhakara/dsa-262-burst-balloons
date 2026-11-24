package main

import "testing"

func TestMaxCoins(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{[]int{3, 1, 5, 8}, 167},
		{[]int{1, 5}, 10},
		{[]int{7}, 7},
		{[]int{}, 0},
		{[]int{2, 3}, 9}, // burst 2 first (2*3=6), then 3 alone (3): 9
	}
	for _, c := range cases {
		if got := MaxCoins(c.in); got != c.want {
			t.Errorf("MaxCoins(%v) = %d, want %d", c.in, got, c.want)
		}
	}
}

// brute tries every burst order by simulation.
func brute(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	best := 0
	for i := 0; i < len(nums); i++ {
		left, right := 1, 1
		if i > 0 {
			left = nums[i-1]
		}
		if i < len(nums)-1 {
			right = nums[i+1]
		}
		rest := append(append([]int{}, nums[:i]...), nums[i+1:]...)
		if v := left*nums[i]*right + brute(rest); v > best {
			best = v
		}
	}
	return best
}

func TestAgainstBrute(t *testing.T) {
	inputs := [][]int{
		{9, 76, 64, 21},
		{5, 4, 3, 2, 1},
		{8, 2, 6, 8, 9},
		{1, 0, 2},
	}
	for _, in := range inputs {
		if got, want := MaxCoins(in), brute(in); got != want {
			t.Errorf("MaxCoins(%v) = %d, brute = %d", in, got, want)
		}
	}
}
