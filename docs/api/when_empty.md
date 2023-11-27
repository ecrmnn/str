# WhenEmpty `*Str`

`WhenEmpty` invokes the given closure if the current string is empty and returns a pointer to the current instance of Str.

### Signature

```go
func (s *Str) WhenEmpty(callback func(s *Str) *Str) *Str
```

### Examples

```go
str.New("go").
	WhenEmpty(func(s *str.Str) *str.Str {
		return s.Append("lang")
	}).
	String()
// "go"

str.New("").
	WhenEmpty(func(s *str.Str) *str.Str {
		return s.Append("golang")
	}).
	String()
// "golang"

```
