package main

func bulkSend(numMessages int) float64 {
	result := 0.0
	for i := range numMessages {
		result += 1.0 + 0.01*float64(i)
	}

	return result
}
