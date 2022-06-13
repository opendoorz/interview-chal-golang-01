package main

import (
	"fmt"
)

func main() {

	// batch is an instance of widgets
	batch := widgets{
		widget{
			Serial: 78,
			SKU:    "Z9000",
		},
		widget{
			Serial: 105,
			SKU:    "GG5000",
		},
		widget{
			Serial: 18,
			SKU:    "ABC123",
		},
		widget{
			Serial: 77,
			SKU:    "Z9000",
		},
		widget{
			Serial: 101,
			SKU:    "GG5000",
		},
		widget{
			Serial: 11,
			SKU:    "ABC123",
		},
		widget{
			Serial: 102,
			SKU:    "GG5000",
		},
	}

	/// 1.1 -- The following line currently initializes the skus variable to nil.  Find and implement the widgets.mapToSKU method below to fix that.
	skus := batch.mapToSKU()

	/// 1.2 -- Sort the widget SKUs (strings) with a one-line call to the sort package on the following line
	//TODO

	/// 1.3 -- Print the widget SKUs sorted alphabetically and with a comma separator by modifying the following line
	fmt.Println("sorted SKUs:", skus)

	// 2 -- Implement the removeDupes function below.  Modify the following line to call removeDupes and print the results with a comma separator.
	fmt.Println("unique sorted SKUs:", skus)

	/// 3.1 -- Find and implement the widgets.sort method below so that the widget objects are properly sorted
	batch.sort()

	/// 3.2 -- Modify the following line to print the sorted widgets as JSON with a 2-space indentation.  Add more lines as needed.
	fmt.Println("sorted widgets:", batch)
}

type widget struct {
	SKU    string `json:"sku"`
	Serial int    `json:"serial"`
}

type widgets []widget

//mapToSKU returns a slice of SKU's, one for each widget in widgets
func (w widgets) mapToSKU() []string {
	//TODO
	return nil
}

//sort orders a slice of widgets by alphabetical SKU, if 2 SKU's are equal widgets are sorted by serial number
func (w widgets) sort() {
	//TODO
}

//removeDupes returns unique string values from input slice
func removeDupes(strs []string) []string {
	return strs
}
