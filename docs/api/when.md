# When `*Str`

`When` invokes the given callback if the given condition is true and returns a pointer to the current instance of Str.

### Signature

```go
func (s *Str) When(condition bool, callback func(s *Str) *Str) *Str
```

### Examples

```go
str.New("go").
	When(true, func(s *str.Str) *str.Str {
		return s.Append("lang")
	}).
	String()
// "golang"

str.New("go").
	When(false, func(s *str.Str) *str.Str {
		return s.Append("lang")
	}).
	String()
// "go"

str.New("").
	When(false, func(s *str.Str) *str.Str {
		return s.Append("lang")
	}).
	String()
// ""

```
