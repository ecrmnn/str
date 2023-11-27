# After `*Str`

`After` returns a new instance of Str containing the substring after the first occurrence of the given substring.

### Signature

```go
func (s *Str) After(substring string) *Str
```

### Examples

```go
str.New("golang").After("go").String()
// "lang"

str.New("go go go!").After("go").String()
// " go go!"

str.New("øæå").After("æ").String()
// "å"

str.New("rocket 🚀🚀 man").After("🚀").String()
// "🚀 man"

```
