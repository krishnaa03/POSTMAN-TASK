package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/xuri/excelize/v2"
)

// Function to get menu items based on day and meal
func getMenuItems(day, meal string) []string {
	// Open the Excel file
	xlsxFile, err := excelize.OpenFile("gomeal.xlsx")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Initialize a variable to store menu items
	var items []string

	// Get all the rows in the Excel file
	rows, err := xlsxFile.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Iterate over each row in the Excel file
	for _, row := range rows {
		// Check if the row contains data for the specified day and meal
		if strings.ToLower(row[0]) == strings.ToLower(day) && strings.ToLower(row[1]) == strings.ToLower(meal) {
			items = append(items, row[2])
		}
	}

	return items
}

func main() {
	// Example usage of the getMenuItems function
	var day string
	fmt.Println("Enter the day: ")
	fmt.Scanln(&day)
	var meal string
	fmt.Println("Enter the meal: ")
	fmt.Scanln(&meal)
	items := getMenuItems(day, meal)

	if len(items) == 0 {
		fmt.Println("No items found for the specified day and meal.")
	} else {
		fmt.Printf("Menu for %s %s: %v\n", day, meal, items)
	}
}
