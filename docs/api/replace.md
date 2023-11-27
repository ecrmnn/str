# Replace `*Str`

`Replace` returns a new instance of Str with all occurrences of the given substring replaced with the given replacement string.

### Signature

```go
func (s *Str) Replace(old string, new string) *Str
```

### Examples

```go
str.New("i like c++").Replace("c++", "go").String()
// "i like go"

str.New("rocket 🚀 man").Replace("🚀", "🌎").String()
// "rocket 🌎 man"

str.New("go, go, go!").Replace("go", "Go").String()
// "Go, Go, Go!"

str.New("你好世界").Replace("你好", "再见").String()
// "再见世界"

str.New("你好世界").Replace("你好", "").String()
// "世界"

```
