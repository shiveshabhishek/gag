package gag

import (
	"fmt"
	"testing"
)

func TestAddition(t *testing.T) {
	tests := []struct {
		a, b int
		want int
	}{
		{10, 5, 15},
		{25, 5, 30},
		{30, 15, 45},
		{30, 9, 39},
		{100, 9, 109},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("(%v,%v)", tc.a, tc.b), func(t *testing.T) {
			got := Addition(tc.a, tc.b)
			if got != tc.want {
				t.Fatalf("Addition() = %v; want %v", got, tc.want)
			}
		})

	}
}
