package main

import (
	"reflect"
	"testing"
)

func TestDistinctInWindows(t *testing.T) {
	cases := []struct {
		in   []int
		k    int
		want []int
	}{
		{[]int{1, 2, 1, 3, 4, 2, 3}, 4, []int{3, 4, 4, 3}},
		{[]int{1, 1, 1}, 2, []int{1, 1}},
		{[]int{1, 2, 3}, 3, []int{3}},
		{[]int{4}, 1, []int{1}},
		{[]int{1, 2}, 3, []int{}}, // k larger than the array
		{[]int{1, 2}, 0, []int{}},
	}
	for _, c := range cases {
		if got := DistinctInWindows(c.in, c.k); !reflect.DeepEqual(got, c.want) {
			t.Errorf("DistinctInWindows(%v,%d) = %v, want %v", c.in, c.k, got, c.want)
		}
	}
}
