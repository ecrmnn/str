# Kebab `*Str`

`Kebab` returns a new lowercased instance of Str with all spaces and underscores replaced with hyphens.

### Signature

```go
func (s *Str) Kebab() *Str
```

### Examples

```go
str.New("").Kebab().String()
// ""

str.New("go go go").Kebab().String()
// "go-go-go"

str.New("GoLang").Kebab().String()
// "go-lang"

str.New("goLang").Kebab().String()
// "go-lang"

str.New("✨🔥✨🔥✨").Kebab().String()
// "✨🔥✨🔥✨"

str.New("✨ 🔥 ✨ 🔥 ✨").Kebab().String()
// "✨-🔥-✨-🔥-✨"

```
