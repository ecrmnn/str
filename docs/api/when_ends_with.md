# WhenEndsWith `*Str`

`WhenEndsWith` invokes the given callback if the current string ends with the given value and returns a pointer to the current instance of Str.

### Signature

```go
func (s *Str) WhenEndsWith(value string, callback func(s *Str) *Str) *Str
```

### Examples

```go
str.New("i like go").
	WhenEndsWith("go", func(s *str.Str) *str.Str {
		return s.Append("lang")
	}).
	String()
// "i like golang"

str.New("i like go").
	WhenEndsWith("gox", func(s *str.Str) *str.Str {
		return s.Append("lang")
	}).
	String()
// "i like go"

str.New("").
	WhenEndsWith("go", func(s *str.Str) *str.Str {
		return s.Append("lang")
	}).
	String()
// ""

```
