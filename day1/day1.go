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
