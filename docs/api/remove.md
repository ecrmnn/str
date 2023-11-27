# Remove `*Str`

`Remove` returns a new instance of Str with all occurrences of the given substring removed.

### Signature

```go
func (s *Str) Remove(old string) *Str
```

### Examples

```go
str.New("i like c++").Remove("c++").String()
// "i like "

str.New("rocket 🚀 man").Remove("🚀").String()
// "rocket  man"

str.New("go, go, go!").Remove("go").String()
// ", , !"

str.New("你好世界").Remove("你好").String()
// "世界"

str.New("").Remove("go").String()
// ""

```
