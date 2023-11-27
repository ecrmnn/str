# WhenNotEmpty `*Str`

`WhenNotEmpty` invokes the given closure if the current string is not empty and returns a pointer to the current instance of Str.

### Signature

```go
func (s *Str) WhenNotEmpty(callback func(s *Str) *Str) *Str
```

### Examples

```go
str.New("go").
	WhenNotEmpty(func(s *str.Str) *str.Str {
		return s.Append("lang")
	}).
	String()
// "golang"

str.New("").
	WhenNotEmpty(func(s *str.Str) *str.Str {
		return s.Append("golang")
	}).
	String()
// ""

```
