# Snake `*Str`

`Snake` returns a new lowercased instance of Str with all spaces and hyphens replaced with underscores.

### Signature

```go
func (s *Str) Snake() *Str
```

### Examples

```go
str.New("go go go").Snake().String()
// "go_go_go"

str.New("GoLang").Snake().String()
// "go_lang"

str.New("already_snake_cased").Snake().String()
// "already_snake_cased"

str.New("").Snake().String()
// ""

str.New("goLang").Snake().String()
// "go_lang"

str.New("✨🔥✨🔥✨").Snake().String()
// "✨🔥✨🔥✨"

str.New("✨ 🔥 ✨ 🔥 ✨").Snake().String()
// "✨_🔥_✨_🔥_✨"

```
