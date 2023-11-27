# WhenContainsAll `*Str`

`WhenContainsAll` invokes the given callback if the current string contains all of the given values and returns a pointer to the current instance of Str.

### Signature

```go
func (s *Str) WhenContainsAll(values []string, callback func(s *Str) *Str) *Str
```

### Examples

```go
str.New("i like go not c++").
	WhenContainsAll([]string{"go", "c++"}, func(s *str.Str) *str.Str {
		return s.Append("!")
	}).
	String()
// "i like go not c++!"

str.New("i like go").
	WhenContainsAll([]string{"a", "b", "c"}, func(s *str.Str) *str.Str {
		return s.Append("lang")
	}).
	String()
// "i like go"

str.New("").
	WhenContainsAll([]string{"go", "lang"}, func(s *str.Str) *str.Str {
		return s.Append("lang")
	}).
	String()
// ""

```
