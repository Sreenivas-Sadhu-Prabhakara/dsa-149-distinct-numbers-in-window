package main

import "fmt"

// DistinctInWindows returns the count of distinct values in every window of
// size k. A frequency map is updated incrementally as the window slides —
// one value enters, one leaves — never recounted from scratch.
func DistinctInWindows(nums []int, k int) []int {
	n := len(nums)
	if k <= 0 || k > n {
		return []int{}
	}
	freq := map[int]int{}
	for _, v := range nums[:k] {
		freq[v]++
	}
	out := make([]int, 0, n-k+1)
	out = append(out, len(freq))
	for i := k; i < n; i++ {
		freq[nums[i]]++
		old := nums[i-k]
		freq[old]--
		if freq[old] == 0 {
			delete(freq, old) // len(freq) stays an honest distinct count
		}
		out = append(out, len(freq))
	}
	return out
}

func main() {
	fmt.Println(DistinctInWindows([]int{1, 2, 1, 3, 4, 2, 3}, 4))
}
