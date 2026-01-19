package main

func getMessageCosts(messages []string) []float64 {
	msgs := make([]float64, len(messages))
	for i := range len(messages) {
		msgs[i] = float64(len(messages[i])) * 0.01
	}

	return msgs
}
