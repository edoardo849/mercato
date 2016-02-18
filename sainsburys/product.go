package sainsburys

import (
	"log"
	"regexp"
	"strconv"
	"strings"

	"../scraper"
	"../utilities"
)

// Product is the struct for a single product in the Grocery Store
type Product struct {
	url         string
	Title       string          `json:"title"`
	Size        string          `json:"size"`
	UnitPrice   utilities.Price `json:"unit_price"`
	Description string          `json:"description"`
}

// NewProduct creates a product and fills missing information
// from the Product page
func NewProduct(title string, price string, url string) *Product {
	p := new(Product)
	p.url = strings.TrimSpace(url)

	p.Title = strings.TrimSpace(title)

	pPage := p.fetchPage()

	p.Size = pPage.GetSizeInKB()
	p.Description = p.getDescription(pPage)

	p.UnitPrice = p.getUnitPrice(price)

	return p
}

func (p *Product) getDescription(pPage *scraper.Scraper) string {
	return strings.TrimSpace(pPage.Find(".productText").First().Text())
}

func (p *Product) getUnitPrice(price string) utilities.Price {

	// Compile a Regex that matches all **positive** Floats
	// assuming no negative price will be displayed
	r := regexp.MustCompile(`[-+]?\d*\.\d+|\d+`)
	sPrice := r.FindStringSubmatch(strings.TrimSpace(price))

	// Check if the Regex has found matches
	if len(sPrice) > 0 {

		// Convert the firstly found match to a float64
		fPrice, err := strconv.ParseFloat(sPrice[0], 64)
		if err != nil {
			log.Fatal(err)
		}

		fPrice = utilities.Round(float64(fPrice), .5, 2)

		return utilities.Price(fPrice)
	}
	return utilities.Price(0)

}

func (p *Product) fetchPage() *scraper.Scraper {

	s := scraper.NewScraper(p.url, "utf-8")
	return s
}

// Products is the list of products in the Grocery Store
type Products []Product
