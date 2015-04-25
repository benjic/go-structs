// Copyright 2015 Benjamin Campbell, All rights reserved
// The use of this code has been provisioned under an Apache License that is
// provided within this repository.

// Helpers and operators related to strings of bits
//
// Bitstrings are sequences of 1's and 0's that have some kind of implicit
// meaning. Included are patterns and functions that make dealing with these bit
// strings easier.
package bitstring

// Crates a series of interrelated channels that iteratively assemble all
// combinations of n choose k bits. The channel returned by the function call
// can be accessed concurrently allowing many go routines.
func CombinationIterator(n, k int) chan int64 {

	ch := make(chan int64, 0)

	switch {
	case n < 0 || k < 0 || n < k:
		// Invalid parameters return empty set
		go func() {
			close(ch)
		}()
	case k == 0:
		go func() {
			ch <- 0
			close(ch)
		}()
	case n == k:
		go func() {
			ch <- (1 << uint(k)) - 1
			close(ch)
		}()
	case k == 1:
		// Recursive base case
		go func() {

			for i := 0; i < n; i++ {
				ch <- 1 << uint(i)
			}

			close(ch)
		}()

	case k > 1:
		// Recursion step
		go func() {

			// for all n choose k-1 bits
			for i := range CombinationIterator(n, k-1) {

				// Iterate over all bits in string
				for j := 1; j < n; j++ {

					// Only emit bitstring if new bit j hits no bit in i
					if (1<<uint(j))&uint(i) == 0 {
						ch <- 1<<uint(j) | i
					} else {
						break
					}
				}
			}

			close(ch)
		}()
	}

	return ch
}
