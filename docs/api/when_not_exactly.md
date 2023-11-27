# WhenNotExactly `*Str`

`WhenNotExactly` invokes the given callback if the current string is not exactly the same as the given value and returns a pointer to the current instance of Str.

### Signature

```go
func (s *Str) WhenNotExactly(value string, callback func(s *Str) *Str) *Str
```

### Examples

```go
str.New("go").
	WhenNotExactly("go", func(s *str.Str) *str.Str {
		return s.Append("lang")
	}).
	String()
// "go"

str.New("").
	WhenNotExactly("go", func(s *str.Str) *str.Str {
		return s.Append("golang")
	}).
	String()
// "golang"

str.New("🚀").
	WhenNotExactly("🚀", func(s *str.Str) *str.Str {
		return s.Append("✨")
	}).
	String()
// "🚀"

```
