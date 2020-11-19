package csvtools

import (
	"fmt"
	"strings"
	"testing"
)

// All columns
// var columns = []string {}
var columns = strings.Split("Name,  Date, Apple,Counter , Some Column", ",")

// Csv data
var header_row = []string { "Counter", "Name", "Date", "Random" }
var data_row = []string { "10740728", "Vasya", "2020-11-18 02:31:00.731", "6701235619237.788567" }

func TestFindColumns(t *testing.T) {
	c, _, m := FindColumns(columns, header_row, false)

	for _, name := range c {
		index := IndexByName(name, m)
		if index == -1 {
			t.Errorf("%s: column does not exists\n", name)
			continue
		}
		
		fmt.Printf("✓ %s: %s\n", name, data_row[index])
	}
	
}
