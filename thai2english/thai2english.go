package thai2english

import (
	"../scrape"
	"../words"
	"fmt"
	"net/url"
)

type Result struct {
	Id           string
	Translations []scrape.TextResult
}

const (
	queryBase   = "http://www.thai2english.com/search.aspx?q="
	soundBase   = "http://www.thai2english.com/sounds/%v.mp3"
	scraperFile = "thai2english/thai2english.xml"
)

var (
	scraper = scrape.Scraper(scraperFile)
)

func Search(text string) words.Word {
	result := fetchResult(text)
	return words.Word{
		Sounds:       soundsFrom(result),
		Translations: scrape.Strings(result.Translations),
	}
}

func soundsFrom(result Result) []string {
	return []string{soundUrl(result.Id)}
}

func fetchResult(text string) Result {
	query := queryUrl(text)
	var result Result
	scraper.ExtractFromURL(query, &result)
	return result
}

func queryUrl(query string) string {
	escapedQuery := url.QueryEscape(query)
	return queryBase + escapedQuery
}

func soundUrl(id string) string {
	return fmt.Sprintf(soundBase, id)
}
