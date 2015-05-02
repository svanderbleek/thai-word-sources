package words

type Word struct {
	Text         string
	Sounds       []string
	Translations []string
}

func (word *Word) Join(other Word) {
	word.Sounds = append(word.Sounds, other.Sounds...)
	word.Translations = append(word.Translations, other.Translations...)
}
