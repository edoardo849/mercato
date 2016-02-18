Mercato: a Sainsbury's Scraper in Go
=====

## Logic
Mercato is a console application intended to scrape a Saynsbury's test page and grab every product's url children page for additional information.

This task could be completed with PHP, Java or Go. However, I have chosen Go because of the possibility to run every scrape request as a concurrent goroutine, thus augmenting the response time.

### Notes

#### GoRoutines order
When we are scraping the store's page, the order of each product in the final JSON object *will not* be the same as the order of appearance on the page. This is because of the GoRoutine's concurrency: the actual order will probably be determined by the responsiveness of the queried page.

If the order of the products in the final JSON object should be the same as the one of the grocery store page, it will be necessary to add an `order` param on the struct `sainsburys.Product` and then implement an order `func` afterwards.

Because of time restraints, this has been left out.

## Dependencies
- `go` installed and `GOPATH` variable well set up. For more info visit  [Go's documentation](https://golang.org/doc/code.html#GOPATH)

### Installation

```bash
# Clone the repo
git clone git@github.com:edoardo849/mercato.git

# Enter the directory
cd mercato

# Install the go dependencies
go get

# Run all the tests
go test ./...

# Run the scraper
go run main.go

```

## Example output from the console

```json
{
   "results":[
      {
         "title":"Sainsbury's Apricot Ripe & Ready x5",
         "size":"38.27kb",
         "unit_price":3.50,
         "description":"Buy Sainsbury's Apricot Ripe & Ready x5 online from Sainsbury's, the same great quality, freshness and choice you'd find in store. Choose from 1 hour delivery slots and collect Nectar points."
      },
      {
         "title":"Sainsbury's Golden Kiwi x4",
         "size":"38.56kb",
         "unit_price":1.80,
         "description":"Buy Sainsbury's Golden Kiwi x4 online from Sainsbury's, the same great quality, freshness and choice you'd find in store. Choose from 1 hour delivery slots and collect Nectar points."
      },
      {
         "title":"Sainsbury's Kiwi Fruit, Ripe & Ready x4",
         "size":"38.98kb",
         "unit_price":1.80,
         "description":"Buy Sainsbury's Kiwi Fruit, Ripe & Ready x4 online from Sainsbury's, the same great quality, freshness and choice you'd find in store. Choose from 1 hour delivery slots and collect Nectar points."
      },
      {
         "title":"Sainsbury's Conference Pears, Ripe & Ready x4 (minimum)",
         "size":"38.54kb",
         "unit_price":1.50,
         "description":"Buy Sainsbury's Conference Pears, Ripe & Ready x4 (minimum) online from Sainsbury's, the same great quality, freshness and choice you'd find in store. Choose from 1 hour delivery slots and collect Nectar points."
      },
      {
         "title":"Sainsbury's Avocado, Ripe & Ready x2",
         "size":"43.44kb",
         "unit_price":1.80,
         "description":"Burgers are a summer must-have and these homemade ones are perfect for a barbecue, topped with cool avocado and served with oven-baked potato wedges."
      },
      {
         "title":"Sainsbury's Avocados, Ripe & Ready x4",
         "size":"38.68kb",
         "unit_price":3.20,
         "description":"Buy Sainsbury's Avocados, Ripe & Ready x4 online from Sainsbury's, the same great quality, freshness and choice you'd find in store. Choose from 1 hour delivery slots and collect Nectar points."
      },
      {
         "title":"Sainsbury's Avocado Ripe & Ready XL Loose 300g",
         "size":"38.67kb",
         "unit_price":1.50,
         "description":"Buy Sainsbury's Avocado Ripe & Ready XL Loose 300g online from Sainsbury's, the same great quality, freshness and choice you'd find in store. Choose from 1 hour delivery slots and collect Nectar points."
      }
   ],
   "total":15.10
}
```
