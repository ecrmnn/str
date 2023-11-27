# Lower `*Str`

`Lower` returns a new instance of Str with all characters in lowercase.

### Signature

```go
func (s *Str) Lower() *Str
```

### Examples

```go
str.New("").Lower().String()
// ""

str.New("✨").Lower().String()
// "✨"

str.New("&").Lower().String()
// "&"

str.New("Æ").Lower().String()
// "æ"

str.New("TEST").Lower().String()
// "test"

```
