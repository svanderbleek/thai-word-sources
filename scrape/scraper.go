package scrape

import (
	"fmt"
	"github.com/ernesto-jimenez/scraperboard"
)

type TextResult struct {
	Text string
}

func Scraper(file string) scraperboard.Scraper {
	scraper, err := scraperboard.NewScraperFromFile(file)
	if err != nil {
		fmt.Println(err)
	}
	return scraper
}

func Strings(results []TextResult) []string {
	var strings []string
	for _, result := range results {
		strings = append(strings, result.Text)
	}
	return strings
}
