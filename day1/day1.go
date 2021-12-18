package day1

func CalculateIncreases(depthReport []int) int {
	var increases int

	for i := 1; i < len(depthReport); i++ {
		if depthReport[i-1] < depthReport[i] {
			increases++
		}
	}

	return increases
}

func CalculateIncreasesWithSlidingWindow(depthReport []int) int {
	var slidingWindowTotals []int

	for i := 1; i < len(depthReport)-1; i++ {
		first := depthReport[i-1]
		second := depthReport[i]
		third := depthReport[i+1]
		slidingWindowTotals = append(slidingWindowTotals, first+second+third)
	}

	return CalculateIncreases(slidingWindowTotals)
}
