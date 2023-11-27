# WhenContainsAny `*Str`

`WhenContainsAny` invokes the given callback if the current string contains any of the given values and returns a pointer to the current instance of Str.

### Signature

```go
func (s *Str) WhenContainsAny(values []string, callback func(s *Str) *Str) *Str
```

### Examples

```go
str.New("i like go").
	WhenContainsAny([]string{"go", "c++"}, func(s *str.Str) *str.Str {
		return s.Append("lang")
	}).
	String()
// "i like golang"

str.New("i like go").
	WhenContainsAny([]string{"a", "b", "c"}, func(s *str.Str) *str.Str {
		return s.Append("lang")
	}).
	String()
// "i like go"

str.New("").
	WhenContainsAny([]string{"go", "lang"}, func(s *str.Str) *str.Str {
		return s.Append("lang")
	}).
	String()
// ""

```
