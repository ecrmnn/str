# AfterLast `*Str`

`AfterLast` returns a new instance of Str containing the substring after the last occurrence of the given substring.

### Signature

```go
func (s *Str) AfterLast(substring string) *Str
```

### Examples

```go
str.New("golang").AfterLast("go").String()
// "lang"

str.New("go go go!").AfterLast("go").String()
// "!"

str.New("øæå").AfterLast("ø").String()
// "æå"

str.New("rocket 🚀🚀 man").AfterLast("🚀").String()
// " man"

```
