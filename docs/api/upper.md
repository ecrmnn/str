# Upper `*Str`

`Upper` returns a new instance of Str with all characters in uppercase.

### Signature

```go
func (s *Str) Upper() *Str
```

### Examples

```go
str.New("").Upper().String()
// ""

str.New("✨").Upper().String()
// "✨"

str.New("&").Upper().String()
// "&"

str.New("æ").Upper().String()
// "Æ"

str.New("test").Upper().String()
// "TEST"

```
