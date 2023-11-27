# Title `*Str`

`Title` returns a new instance of Str with the first letter of each word capitalized.

### Signature

```go
func (s *Str) Title() *Str
```

### Examples

```go
str.New("go øl åre ære").Title().String()
// "Go Øl Åre Ære"

str.New("go ✨ go").Title().String()
// "Go ✨ Go"

str.New("_ _ _").Title().String()
// "_ _ _"

str.New("").Title().String()
// ""

```
