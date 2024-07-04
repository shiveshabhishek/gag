package gag

import (
	"fmt"
	"testing"
)

func TestMin(t *testing.T) {
	tests := []struct {
		a, b int
		want int
	}{
		{12, 3, 3},
		{44, 55, 44},
		{23, 43, 23},
		{-1, 5, -1},
		{42, 21, 21},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("(%v,%v)", tc.a, tc.b), func(t *testing.T) {
			got := Min(tc.a, tc.b)
			if got != tc.want {
				t.Fatalf("Min() = %v; want %v", got, tc.want)
			}

		})
	}

}

// another way for testing
func TestMax(t *testing.T) {
	if got := Max(3, 2); got != 3 {
		t.Errorf("Max(3, 2) = %d; want 1", got)
	}
	if got := Max(2, 3); got != 3 {
		t.Errorf("Max(2, 3) = %d; want 3", got)
	}
	if got := Max(99, 100); got != 100 {
		t.Errorf("Max(99, 100) = %d; want 100", got)
	}
}
