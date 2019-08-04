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
	flag.Parse()
	numStreams := *numStreamsPtr

	streams := crawl.Scrape(numStreams)

	records := make([][]string, numStreams+1)
	records[0] = crawl.Fields()
	for i := 0; i < numStreams; i++ {
		records[i+1] = streams[i].Slice()
	}

	f, err := os.Create("streams.csv")
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
