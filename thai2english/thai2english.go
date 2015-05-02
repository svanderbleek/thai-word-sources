package thai2english

import (
	"../words"
	"fmt"
	"github.com/ernesto-jimenez/scraperboard"
	"net/url"
)

func Search(text string) words.Word {
	result := fetchResult(text)
	fmt.Println(result)
	return words.Word{
		Sounds:       soundsFrom(result),
		Translations: translationsFrom(result),
	}
}

func soundsFrom(result Result) []string {
	return []string{soundUrl(result.Id)}
}

func translationsFrom(result Result) []string {
	var translations []string
	for _, translation := range result.Translations {
		translations = append(translations, translation.Translation)
	}
	return translations
}

const (
	queryBase   = "http://www.thai2english.com/search.aspx?q="
	soundBase   = "http://www.thai2english.com/sounds/%v.mp3"
	scraperFile = "thai2english/thai2english.xml"
)

type Translation struct {
	Translation string
}

type Result struct {
	Id           string
	Translations []Translation
}

var (
	scraper = buildScraper()
)

func buildScraper() scraperboard.Scraper {
	scraper, err := scraperboard.NewScraperFromFile(scraperFile)
	if err != nil {
		panic(err)
	}
	return scraper
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
