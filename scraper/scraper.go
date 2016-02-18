package scraper

import (
	"log"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/djimenez/iconv-go"
)

// Scraper is an alias for goquery.Document. Reason for this
// is that Go cannot extend existing types in another package
type Scraper struct {
	url           string
	charset       string
	doc           *goquery.Document
	contentLength int64
}

// NewScraper is a wrapper intended to be used with *goquery.Document:
// it fetches a remote url, converts the content to the specified charset
// and calculates the page size
func NewScraper(url string, charset string) *Scraper {
	s := new(Scraper)
	s.url = url
	s.charset = charset
	s.doc = s.getDocument()
	s.contentLength = s.getContentLength()

	// Setup and return our Document struct
	return s
}

// GetSizeInKB converts the length of the document in KB
func (s *Scraper) GetSizeInKB() string {
	sizeInKb := float64(s.contentLength) / 1024

	return strconv.FormatFloat(sizeInKb, 'f', 2, 64) + "kb"
}

// Find is just a wrapper for *goquery.Document().Find()
// as this is part of an already tested library, we will
// not build unit tests on this func
func (s *Scraper) Find(selector string) *goquery.Selection {
	return s.doc.Find(selector)
}

func (s *Scraper) getContentLength() int64 {

	var l int64

	// If the headers have a Content Length, we will accept it,
	// https://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html
	// otherwise we will calculate it
	if s.contentLength > 0 {

		l = s.contentLength

	} else {
		// get the raw html
		html, err := s.doc.Html()
		if err != nil {
			log.Fatal(err)
		}
		l = int64(len(html))
	}

	return l
}

func (s *Scraper) getDocument() *goquery.Document {

	res := s.getResponse()
	defer res.Body.Close()

	// Convert the designated charset HTML to utf-8 encoded HTML.
	// `charset` being one of the charsets known by the iconv package.
	utfBody, err := iconv.NewReader(res.Body, s.charset, "utf-8")
	if err != nil {
		log.Fatal(err)
	}

	// use utfBody using goquery
	doc, err := goquery.NewDocumentFromReader(utfBody)
	if err != nil {
		log.Fatal(err)
	}
	return doc
}

func (s *Scraper) getResponse() *http.Response {
	// Load the URL
	res, err := http.Get(s.url)
	if err != nil {
		log.Fatal(err)
	}
	// Get the Header's content Length
	s.contentLength = res.ContentLength

	return res
}
