# WhenExactly `*Str`

`WhenExactly` invokes the given callback if the current string is exactly the same as the given value and returns a pointer to the current instance of Str.

### Signature

```go
func (s *Str) WhenExactly(value string, callback func(s *Str) *Str) *Str
```

### Examples

```go
str.New("go").
	WhenExactly("go", func(s *str.Str) *str.Str {
		return s.Append("lang")
	}).
	String()
// "golang"

str.New("").
	WhenExactly("go", func(s *str.Str) *str.Str {
		return s.Append("golang")
	}).
	String()
// ""

str.New("🚀").
	WhenExactly("🚀", func(s *str.Str) *str.Str {
		return s.Append("✨")
	}).
	String()
// "🚀✨"

```
