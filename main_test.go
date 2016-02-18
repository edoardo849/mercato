package main

import (
	"encoding/json"
	"testing"

	"./sainsburys"
)

func TestGroceryStore(t *testing.T) {
	s := GroceryStore()
	bytes := []byte(s)
	// Unmarshal string into structs.
	var groceryStore sainsburys.GroceryStore
	json.Unmarshal(bytes, &groceryStore)

	if groceryStore.Total <= 0 {
		t.Error("The total should be more than 0")
	}

	if len(groceryStore.Products) == 0 {
		t.Error("The product struct should not be empty")
	}
}
