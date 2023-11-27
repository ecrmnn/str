# Last `*Str`

`Last` returns s new instance of Str containing the last character of the current string.

### Signature

```go
func (s *Str) Last() *Str
```

### Examples

```go
str.New("golang").Last().String()
// "g"

str.New("hello world").Last().String()
// "d"

str.New("🔥 Go 🤩✨").Last().String()
// "✨"

str.New("æøå").Last().String()
// "å"

str.New("").Last().String()
// ""

```
