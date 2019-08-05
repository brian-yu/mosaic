package main

import (
	"encoding/csv"
	"flag"
	"github.com/brian-yu/mosaic/pkg/crawl"
	"log"
	"os"
)

func main() {

	numStreamsPtr := flag.Int("limit", 50, "number of streams to scrape")
	fname := flag.String("o", "streams.csv", "file to write streams to")
	flag.Parse()
	numStreams := *numStreamsPtr

	streams := crawl.ScrapePopular(numStreams)

	records := make([][]string, numStreams+1)
	records[0] = crawl.Fields()
	for i := 0; i < numStreams; i++ {
		records[i+1] = streams[i].Slice()
	}

	writeCsv(*fname, records)
}

func writeCsv(fname string, records [][]string) {
	f, err := os.Create(fname)
	if err != nil {
		log.Fatal(err)
	}

	w := csv.NewWriter(f)

	for _, record := range records {
		if err := w.Write(record); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}

	// Write any buffered data to the underlying writer (standard output).
	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}
