package crawl

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const RawInsecamUrl = "http://www.insecam.org/en/byrating/"

// Scrapes numStreams streams from Insecam popular.
func Scrape(numStreams int) []*Stream {
	streams := make([]*Stream, numStreams)
	numScraped := 0
	page := 1
	for numScraped < numStreams {
		scrapePage(page, &numScraped, numStreams, streams)
		page++
	}
	return streams
}

// Scrapes streams on a specified page. Will scrape all streams unless count
// exceeds limit. Streams are placed in the streams slice.
func scrapePage(page int, count *int, limit int, streams []*Stream) {
	// Construct the URL.
	u, err := url.Parse(RawInsecamUrl)
	if err != nil {
		log.Fatal(err)
	}
	q := u.Query()
	q.Set("page", strconv.Itoa(page))
	u.RawQuery = q.Encode()

	// Send a GET request.
	res, err := http.Get(u.String())
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document.
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the stream items.
	doc.Find(".thumbnail-item").Each(func(i int, s *goquery.Selection) {
		// Stop after we have enough streams.
		if *count >= limit {
			return
		}
		streams[*count] = constructStream(s)
		fmt.Printf("Scraping stream %d: %s\n", *count, streams[*count].String())
		*count++
	})
}

// Parses img src and title into a Stream struct.
func constructStream(s *goquery.Selection) *Stream {
	src, _ := s.Find("img").Attr("src")
	title, _ := s.Find("img").Attr("title")
	location := extractLocation(title)
	relUrl, _ := s.Find("a").Attr("href")
	insecamId := strings.Split(relUrl, "/")[3]
	return &Stream{src: src, location: location, insecamId: insecamId}
}

// Extracts location from title string.
func extractLocation(title string) string {
	countryState := strings.Split(
		strings.Split(title, " in ")[1],
		", ",
	)
	return countryState[1] + ", " + countryState[0]
}
