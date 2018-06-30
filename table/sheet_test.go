package main

import "testing"

func TestSheet(t *testing.T) {
	values := [][]interface{}{{"1", "2"}, {"3", "4"}}
	createSheetAndUpdate(values)
}
