package main

func getNameCounts(names []string) map[rune]map[string]int {
	m := map[rune]map[string]int{}
	for _, name := range names {
		if len(name) == 0 {
			continue
		}
		runes := []rune(name)
		initial := runes[0]
		if m[initial] == nil {
			m[initial] = make(map[string]int)
		}
		m[initial][name]++
	}

	return m
}
