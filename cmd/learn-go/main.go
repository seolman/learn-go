package main

func countReports(numSentCh chan int) int {
	total := 0
	for num := range numSentCh {
		total += num
	}
	return total
}

// don't touch below this line

func sendReports(numBatches int, ch chan int) {
	for i := range numBatches {
		numReports := i*23 + 32%17
		ch <- numReports
	}
	close(ch)
}
