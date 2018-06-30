package main

import "strings"

func main() {
	table := getDirectReport(true)

	values := [][]interface{}{}

	for _, line := range strings.Split(strings.TrimSuffix(table, "\n"), "\n") {
		line2 := []interface{}{}
		for _, value := range strings.Split(line, "\t") {
			line2 = append(line2, interface{}(value))
		}
		values = append(values, line2)
	}

	createSheetAndUpdate(values)
}
