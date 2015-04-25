// Copyright 2015 Benjamin Campbell, All rights reserved
// The use of this code has been provisioned under an Apache License that is
// provided within this repository.

package bitstring

import (
	"fmt"
	"strconv"
	"testing"
)

// Compare bitstring outputs against hand testing combinations
func TestBitStringCombinations(t *testing.T) {

	cases := []struct {
		n, k int
		want []int64
	}{
		{-3, 1, []int64{}},
		{3, -1, []int64{}},
		{3, 0, []int64{0}},
		{3, 1, []int64{1, 2, 4}},
		{3, 2, []int64{3, 5, 6}},
		{3, 3, []int64{7}},
		{4, 0, []int64{0}},
		{4, 1, []int64{1, 2, 4, 8}},
		{4, 2, []int64{3, 5, 9, 6, 10, 12}},
		{4, 3, []int64{7, 11, 13, 14}},
		{4, 4, []int64{15}},
		{5, 0, []int64{0}},
		{5, 1, []int64{1, 2, 4, 8, 16}},
		{5, 2, []int64{3, 5, 9, 17, 6, 10, 12, 18, 20, 24}},
		{5, 3, []int64{7, 11, 13, 19, 21, 25, 14, 22, 26, 28}},
		{5, 4, []int64{15, 23, 27, 29, 30}},
		{5, 5, []int64{31}},
	}

	for _, c := range cases {

		got := make([]int64, 0)

		for i := range CombinationIterator(c.n, c.k) {
			got = append(got, i)
		}

		if len(got) != len(c.want) {
			t.Errorf("Expected %d permutations, got %d %+v %+v", len(c.want), len(got), c.want, got)
		} else {
			for i := range got {
				if got[i] != c.want[i] {
					t.Errorf("Different permutatation: want=%+v, got=%+v", c.want, got)
				}
			}
		}
	}
}

// Compare the number of combinations for n choose k combinations
func TestCombinationCount(t *testing.T) {

	cases := []struct {
		n, k, want int
	}{
		{12, 6, 924},
	}

	for _, c := range cases {
		got := 0
		for range CombinationIterator(c.n, c.k) {
			got++
		}

		if got != c.want {
			t.Errorf("%d choose %d has count %d, got %d", c.n, c.k, c.want, got)
		}
	}

}

// We can generate all combinations of choosing one bit from 3 possible bits.
func ExampleCombinationIterator() {
	for c := range CombinationIterator(3, 1) {
		fmt.Printf("%03s\n", strconv.FormatInt(c, 2))
	}
	// Output:
	// 001
	// 010
	// 100
}
