package main

func indexOfFirstBadWord(msgs []string, badWords []string) int {
	for i, msg := range msgs {
		for _, bad := range badWords {
			if msg == bad {
				return i
			}
		}
	}
	return -1
}
