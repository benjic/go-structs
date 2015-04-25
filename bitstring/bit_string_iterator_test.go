package bitstring

import "testing"

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

		for i := range Iterator(c.n, c.k) {
			got = append(got, i)
		}

		if len(got) != len(c.want) {
			t.Errorf("Expected %d permutations, got %d %+v %+v", len(c.want), len(got), c.want, got)
		} else {
			for i := range got {
				if got[i] != c.want[i] {
					t.Errorf("Differnt permutatation: want=%+v, got=%+v", c.want, got)
				}
			}
		}
	}
}

func TestCombinationCount(t *testing.T) {

	cases := []struct {
		n, k, want int
	}{
		{12, 6, 924},
	}

	for _, c := range cases {
		got := 0
		for range Iterator(c.n, c.k) {
			got++
		}

		if got != c.want {
			t.Error("%d choose %d has count %d, got %d", c.n, c.k, c.want, got)
		}
	}

}
