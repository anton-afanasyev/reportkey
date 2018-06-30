package main

import (
	"fmt"
	"testing"
)

func TestDirect(t *testing.T) {
	table := getDirectReport(true)
	fmt.Println("Table: \n", table)
}
