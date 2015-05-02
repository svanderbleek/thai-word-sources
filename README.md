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
  "sources"
  "thai2english"
  "forvo"
)

source := sources.Bundle(thai2english.Search, forvo.Search)
word := source("ไม่")
```

There is a server that demos functionality.

```
go build -o server
./server
open http://localhost:8080/
```
