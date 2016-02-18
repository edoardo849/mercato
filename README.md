Mercato: a Sainsbury's Scraper in Go
=====

## Logic
Mercato is a console application intended to scrape a Saynsbury's test page and grab every product's url children page for additional information.

This task could be completed with PHP, Java or Go. However, I have chosen Go because of the possibility to run every scrape request as a concurrent goroutine, thus augmenting the response time.

### Notes

#### Description on the product page
When I am required to

>follow each link and get the size (in kb) of the linked HTML (no assets)
and the description to display in the JSON

I see on the product's page that there is just a short text in the description that does not correspond to the example. If I check the source code though, the ".productText" html class corresponds to many items under the "Information" Tab.

For example: http://hiring-tests.s3-website-eu-west-1.amazonaws.com/2015_Developer_Scrape/sainsburys-avocado--ripe---ready-x2.html
Description: "Avocados" (aka: first instance of `".productText"`)

I have chosen to grab only that description: if it's required to grab the whole "Info" page I can quickly modify the code.

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
