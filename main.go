package main

import (
	"fmt"

	"./sainsburys"
	"./utilities"
)

// this is the console application
func main() {
	fmt.Println(GroceryStore())
}

// GroceryStore is the function that will
// scrape the Sainsbury's Grocery store page
// and return a JSON string of the items
func GroceryStore() string {
	s := sainsburys.NewGroceryStore()
	res2B, _ := utilities.JSONMarshal(s, true)
	return string(res2B)
}
