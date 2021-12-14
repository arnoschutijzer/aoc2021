package day1

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func ReadReports(path string) []int {
	data, _ := ioutil.ReadFile(path)
	stringInput := string(data)
	reportsInStrings := strings.Split(stringInput, "\n")

	var reports []int
	for _, item := range reportsInStrings {
		value, _ := strconv.Atoi(item)
		reports = append(reports, value)
	}

	return reports
}
