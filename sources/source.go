package sources

import (
	"../words"
)

type Source func(string) words.Word

type Bundled struct {
	Sources []Source
}

func (bundle *Bundled) Search(query string) words.Word {
	var word words.Word
	var next words.Word
	for _, source := range bundle.Sources {
		next = source(query)
		word.Join(next)
	}
	return word
}

func Bundle(sources ...Source) Source {
	bundle := Bundled{
		Sources: sources,
	}
	return bundle.Search
}
