package sainsburys

import "testing"

func TestNewGroceryStore(t *testing.T) {

	g := NewGroceryStore()

	if len(g.Products) < 1 {
		t.Error("The product struct should not be empty")
	}

	if g.Total < 1 {
		t.Error("The sum of the prices should not be 0")
	}
}
