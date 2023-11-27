# WhenContains `*Str`

`WhenContains` invokes the given callback if the current string contains the given value and returns a pointer to the current instance of Str.

### Signature

```go
func (s *Str) WhenContains(value string, callback func(s *Str) *Str) *Str
```

### Examples

```go
str.New("i like go").
	WhenContains("go", func(s *str.Str) *str.Str {
		return s.Append("lang")
	}).
	String()
// "i like golang"

str.New("i like go").
	WhenContains("gox", func(s *str.Str) *str.Str {
		return s.Append("lang")
	}).
	String()
// "i like go"

str.New("").
	WhenContains("go", func(s *str.Str) *str.Str {
		return s.Append("lang")
	}).
	String()
// ""

```
