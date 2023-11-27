# Before `*Str`

`Before` returns a new instance of Str containing the substring before the first occurrence of the given substring.

### Signature

```go
func (s *Str) Before(substring string) *Str
```

### Examples

```go
str.New("golang").Before("go").String()
// ""

str.New("go go go!").Before("go!").String()
// "go go "

str.New("øæå").Before("æ").String()
// "ø"

str.New("rocket 🚀🚀 man").Before("🚀").String()
// "rocket "

str.New("aaa").Before("a").String()
// ""

str.New("aabbcc").Before("bc").String()
// "aab"

```
