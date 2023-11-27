# Pascal `*Str`

`Pascal` returns a new instance of Str in PascalCase.

### Signature

```go
func (s *Str) Pascal() *Str
```

### Examples

```go
str.New("").Pascal().String()
// ""

str.New("go go go").Pascal().String()
// "GoGoGo"

str.New("hello world").Pascal().String()
// "HelloWorld"

str.New("Go-lang").Pascal().String()
// "Go-lang"

str.New("goLang").Pascal().String()
// "GoLang"

str.New("✨🔥✨🔥✨").Pascal().String()
// "✨🔥✨🔥✨"

str.New("✨ 🔥 ✨ 🔥 ✨").Pascal().String()
// "✨🔥✨🔥✨"

```
