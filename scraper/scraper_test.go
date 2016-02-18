package scraper

import (
	"reflect"
	"regexp"
	"strings"
	"testing"
)

var urlTests = []struct {
	url     string // input
	charset string // expected result
}{
	{"http://icomarini.com/blog", "utf-8"},
	{"http://www.bbc.co.uk/programmes/p02nq0gn/episodes/downloads", "utf-8"},
	{"http://www.repubblica.it/", "utf-8"},
	{"http://hiring-tests.s3-website-eu-west-1.amazonaws.com/2015_Developer_Scrape/5_products.html", "utf-8"},
}

func TestNewGetScraper(t *testing.T) {

	for _, tt := range urlTests {
		t.Logf("Testing Scraper for %s", tt.url)

		actual := NewScraper(tt.url, tt.charset)

		if actual.charset != tt.charset {
			t.Errorf("Url(%s): expected charset %s, actual %s", tt.url, tt.charset, actual.charset)
		}

		expType := "*goquery.Document"
		actualType := reflect.TypeOf(actual.doc).String()
		if actualType != expType {
			t.Errorf("Url(%s): expected type %s, actual %s", tt.url, expType, actualType)
		}

		if actual.contentLength <= 0 {
			t.Errorf("Url(%s): expected size >0 , actual %x", tt.url, actual.contentLength)
		}

	}
}

func TestGetSizeInKB(t *testing.T) {
	for _, tt := range urlTests {

		t.Logf("Testing Size For %s", tt.url)
		actual := NewScraper(tt.url, tt.charset).GetSizeInKB()

		if !strings.Contains(actual, "kb") {
			t.Errorf("Url(%s): expected to contain text kb , actual %s", tt.url, actual)
		}

		r, _ := regexp.Compile(`[0-9]`)
		if !r.MatchString(actual) {
			t.Errorf("Url(%s): expected to contain numbers , actual %s", tt.url, actual)
		}

	}
}
