package sainsburys

import "testing"

func TestProduct(t *testing.T) {
	p := NewProduct("  Product Test  ", "€43.6", "http://hiring-tests.s3-website-eu-west-1.amazonaws.com/2015_Developer_Scrape/sainsburys-avocado-xl-pinkerton-loose-300g.html")

	// We are **NOT** testing the size because it is part
	// of another package that has been already testes

	if p.Title != "Product Test" {
		t.Error("The title shoulf be trimmed from trailing spaces")
	}

	if p.UnitPrice < 0 {
		t.Error("The price cannot be a negative number")
	}

	if p.Description != "Buy Sainsbury's Avocado Ripe & Ready XL Loose 300g online from Sainsbury's, the same great quality, freshness and choice you'd find in store. Choose from 1 hour delivery slots and collect Nectar points." {
		t.Error("The description does not correspond to the one at the remote page")
	}

}
