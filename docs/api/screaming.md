# Screaming `*Str`



### Signature

```go
func (s *Str) Screaming() *Str
```

### Examples

```go
str.New("go go go").Screaming().String()
// "GO_GO_GO"

str.New("GoLang").Screaming().String()
// "GO_LANG"

str.New("ALREADY_SCREAMING").Screaming().String()
// "ALREADY_SCREAMING"

str.New("").Screaming().String()
// ""

str.New("goLang").Screaming().String()
// "GO_LANG"

str.New("✨🔥✨🔥✨").Screaming().String()
// "✨🔥✨🔥✨"

str.New("✨ 🔥 ✨ 🔥 ✨").Screaming().String()
// "✨_🔥_✨_🔥_✨"

```
