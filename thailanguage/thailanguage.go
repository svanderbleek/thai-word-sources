package thailanguage

import (
	"../scrape"
	"../words"
	"fmt"
	"net/http"
	"net/url"
)

type Result struct {
	Id           string
	Translations []scrape.TextResult
}

const (
	dictionary  = "http://www.thai-language.com/default.aspx"
	soundBase   = "http://www.thai-language.com/mp3/E%v.mp3"
	scraperFile = "thailanguage/thailanguage.xml"
)

var (
	scraper = scrape.Scraper(scraperFile)
)

func Search(query string) words.Word {
	result := fetchResult(query)
	fmt.Println(result)
	return words.Word{
		Sounds:       soundsFrom(result),
		Translations: scrape.Strings(result.Translations),
	}
}

func soundUrl(id string) string {
	return fmt.Sprintf(soundBase, id)
}

func soundsFrom(result Result) []string {
	return []string{soundUrl(result.Id)}
}

func fetchResult(query string) Result {
	var result Result
	response := queryResponse(query)
	scraper.ExtractFromResponse(response, &result)
	return result
}

func queryResponse(search string) *http.Response {
	query := url.Values{
		"tmode":  {"0"},
		"enmode": {"1"},
		"search": {search},
	}
	response, _ := http.PostForm(dictionary, query)
	return response
}
