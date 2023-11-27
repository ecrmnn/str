# BeforeLast `*Str`

`BeforeLast` returns a new instance of Str containing the substring before the last occurrence of the given substring.

### Signature

```go
func (s *Str) BeforeLast(substring string) *Str
```

### Examples

```go
str.New("golang").BeforeLast("go").String()
// ""

str.New("go go go!").BeforeLast("go").String()
// "go go "

str.New("øæøåøåø").BeforeLast("ø").String()
// "øæøåøå"

str.New("rocket 🚀🚀 man").BeforeLast("🚀").String()
// "rocket 🚀"

```
