package day1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Day1Part1ExampleHas7Increases(t *testing.T) {
	reports := ReadReports("./fixtures/example.txt")
	increases := CalculateIncreases(reports)
	assert.Equal(t, 7, increases)
}

func Test_Part1(t *testing.T) {
	reports := ReadReports("./fixtures/day1.txt")
	increases := CalculateIncreases(reports)
	assert.Equal(t, 1390, increases)
}

func Test_Day1Part2ExampleHas5IncreasesWithSlidingWindow(t *testing.T) {
	reports := ReadReports("./fixtures/example.txt")
	increases := CalculateIncreasesWithSlidingWindow(reports)
	assert.Equal(t, 5, increases)
}

func TestPart2(t *testing.T) {
	reports := ReadReports("./fixtures/day1.txt")
	increases := CalculateIncreasesWithSlidingWindow(reports)
	assert.Equal(t, 1457, increases)
}
