package day1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Day1ExampleHas7Increases(t *testing.T) {
	reports := ReadReports("./fixtures/example.txt")
	increases := CalculateIncreases(reports)
	assert.Equal(t, 7, increases)
}

func Test_Part1(t *testing.T) {
	reports := ReadReports("./fixtures/part1.txt")
	increases := CalculateIncreases(reports)
	assert.Equal(t, 1390, increases)
}
