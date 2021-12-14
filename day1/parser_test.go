package day1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ReadReports(t *testing.T) {
	actual := ReadReports("./fixtures/parser.txt")
	expected := []int{
		100,
		100,
		102,
		103,
		123123,
		3959,
		599544,
	}
	assert.Equal(t, actual, expected)
}
