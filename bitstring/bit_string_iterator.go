package bitstring

func CombinationIterator(n, k int) chan int64 {

	ch := make(chan int64, 0)

	switch {
	case n < 0 || k < 0 || n < k:
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
		go func() {

			for i := 0; i < n; i++ {
				ch <- 1 << uint(i)
			}

			close(ch)
		}()

	case k > 1:
		go func() {

			for i := range CombinationIterator(n, k-1) {

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
