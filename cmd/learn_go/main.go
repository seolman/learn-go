package main

type cost struct {
	day   int
	value float64
}

func getDayCosts(costs []cost, day int) []float64 {
	dayCosts := []float64{}
	for i := range costs {
		cost := costs[i]
		if cost.day == day {
			dayCosts = append(dayCosts, cost.value)
		}
	}
	return dayCosts
}
