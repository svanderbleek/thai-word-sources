## Word Source

Go packages for looking up Thai words. Translation and Audio.

```go
type Word struct {
  Text         string
  Translations []string
  Sounds       []string
}
```

Translations are text string and sounds are url strings.
There are multiple word sources.

```
import (
  "thai2english"
  "forvo"
)

th2en_word := thai2english.Search("ไม่")
forvo_word := forvo.Search("ไม่")
```

Results can be combined from sources.

```
import (
  "words"
  "thai2english"
  "forvo"
)

source := words.CombineSources(thai2english, forvo)
word := source.Search("ไม่")
```
