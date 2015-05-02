package forvo

import (
	"../words"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

const (
	queryBase = "http://apifree.forvo.com/key/%v/format/xml/action/word-pronunciations/word/%v/language/th"
)

var (
	apiKey = os.Getenv("FORVO_API_KEY")
)

type Item struct {
	Pathmp3 string `xml:"pathmp3"`
}

type Result struct {
	Items []Item `xml:"item"`
}

func Search(query string) words.Word {
	result := fetchResult(query)
	word := words.Word{
		Sounds: soundsFrom(result),
	}
	return word
}

func soundsFrom(result Result) []string {
	var sounds []string
	for _, item := range result.Items {
		sounds = append(sounds, item.Pathmp3)
	}
	return sounds
}

func queryUrl(word string) string {
	query := url.QueryEscape(word)
	return fmt.Sprintf(queryBase, apiKey, query)
}

func fetchResult(query string) Result {
	url := queryUrl(query)
	body := responseBody(url)
	var result Result
	xml.Unmarshal(body, &result)
	return result
}

func responseBody(url string) []byte {
	response, _ := http.Get(url)
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	return body
}
