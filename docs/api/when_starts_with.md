# WhenStartsWith `*Str`

`WhenStartsWith` invokes the given callback if the current string starts with the given value and returns a pointer to the current instance of Str.

### Signature

```go
func (s *Str) WhenStartsWith(value string, callback func(s *Str) *Str) *Str
```

### Examples

```go
str.New("go").
	WhenStartsWith("go", func(s *str.Str) *str.Str {
		return s.Append("lang")
	}).
	String()
// "golang"

str.New("go").
	WhenStartsWith("gox", func(s *str.Str) *str.Str {
		return s.Append("lang")
	}).
	String()
// "go"

str.New("").
	WhenStartsWith("go", func(s *str.Str) *str.Str {
		return s.Append("lang")
	}).
	String()
// ""

```
