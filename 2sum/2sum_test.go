package gag

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		arrayElem []int
		sum       int
		pairIndex []int
	}{
		{[]int{1, 2, 3, 4, 5}, 9, []int{3, 4}},
		{[]int{5, 4, 3, 2, 1}, 5, []int{2, 3}},
		{[]int{33, 2, 2, 11, 34, 5, 6, 7, 56, 87, 8, 0}, 87, []int{9, 11}},
		{[]int{3, 2, 3, 4, 2, 3}, 5, []int{0, 1}},
		{[]int{3, 2, 3, 4, 2, 3}, 4, []int{1, 4}},
		{[]int{3, 2, 3, 4}, 4, nil},
	}
	for _, tc := range tests {
		got := TwoSum(tc.arrayElem, tc.sum)
		if !reflect.DeepEqual(got, tc.pairIndex) {
			t.Errorf("TwoSum(%v, %d) = %v; want %v", tc.arrayElem, tc.sum, got, tc.pairIndex)
		}
	}
}
