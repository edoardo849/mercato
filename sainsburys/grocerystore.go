//Package sainsburys Here we isolate package sainsburys from the rest of the application
package sainsburys

import (
	"sync"

	"../scraper"
	"../utilities"
	"github.com/PuerkitoBio/goquery"
)

// GroceryStore is the model for the store
type GroceryStore struct {
	url      string
	Products []*Product      `json:"results"`
	Total    utilities.Price `json:"total"`
}

// NewGroceryStore Loop through every item with a goRoutine
func NewGroceryStore() *GroceryStore {

	g := new(GroceryStore)
	g.url = "http://hiring-tests.s3-website-eu-west-1.amazonaws.com/2015_Developer_Scrape/5_products.html"
	g.Products = g.getAsyncProducts()

	return g
}

// Get all the products in the page with a goRoutine
func (g *GroceryStore) getAsyncProducts() []*Product {

	// Start by scraping the store page
	s := g.fetchStore()

	// Select all the products in the page
	pSelector := s.Find(".product")

	// Count how many products there are in the page
	pLen := pSelector.Size()

	// This is the slice that will contain all our products
	var p []*Product

	// Construct a slice of product chan, make it
	// of the size of the products found in the page
	pChan := make(chan *Product, pLen)

	g.fetchProducts(pChan, pSelector, pLen)

	// Close the channel when the previous function is done
	close(pChan)

	// This is where we will store the
	// sum of all unit prices
	var total utilities.Price
	total = 0
	for i := range pChan {
		p = append(p, i)
		total += i.UnitPrice
	}

	// Assign the Total to our GroceryStore struct
	g.Total = total

	return p
}

// The function that will launch our GoRoutines: in order to prevent race conditions
// we will sync all of our routines
func (g *GroceryStore) fetchProducts(products chan *Product, pSelector *goquery.Selection, pLen int) {
	var wg sync.WaitGroup

	// We are telling to the WaitGroup
	// how many items should be in the channel
	// before it completes (the number is of course
	// the total of our products)
	wg.Add(pLen)

	// Loop through every product in the page
	pSelector.Each(func(i int, s *goquery.Selection) {

		title := s.Find("h3")
		url, _ := title.Find("a").Attr("href")
		price := s.Find(".pricePerUnit").Text()

		// Fetch every product concurrently with
		// a GoRoutine
		go func(products chan *Product) {
			defer wg.Done()
			products <- NewProduct(title.Text(), price, url)

		}(products)
	})

	wg.Wait()
}

func (g *GroceryStore) fetchStore() *scraper.Scraper {
	s := scraper.NewScraper(g.url, "utf-8")

	return s
}
