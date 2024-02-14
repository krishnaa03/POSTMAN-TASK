package main

import (
	"fmt"
)

// Define a struct to represent each meal
type Meal struct {
	Day   string
	Date  string
	Type  string
	Items []string
}

// Method to print the details of each meal instance
func (m *Meal) PrintDetails() {
	fmt.Printf("Day: %s\n", m.Day)
	fmt.Printf("Date: %s\n", m.Date)
	fmt.Printf("Meal: %s\n", m.Type)
	fmt.Println("Items:")
	for _, item := range m.Items {
		fmt.Println("-", item)
	}
	fmt.Println()
}

func main() {
	// Create an instance of Meal
	meal1 := Meal{
		Day:   "FRIDAY",
		Date:  "2024-02-09",
		Type:  "BREAKFAST",
		Items: []string{"CHOICE OF EGG", "CORNFLAKES", "TEA/COFFEE", "MILK"},
	}

	// Print the details of meal1
	meal1.PrintDetails()

	// Create another instance of Meal
	meal2 := Meal{
		Day:   "Monday",
		Date:  "2024-02-05",
		Type:  "Lunch",
		Items: []string{"DISCO PAPAD", "KADHI PAKODA", "MADRASI ALOO", "VEG KHICHDI", "CHAPATI", "PLAIN RICE"},
	}

	// Print the details of meal2
	meal2.PrintDetails()
}
