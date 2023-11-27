# Contains `bool`

`Contains` returns true if the current string contains the given substring.

### Signature

```go
func (s *Str) Contains(substring string) bool
```

### Examples

```go
str.New("golang").Contains("lan")
// true

str.New("golang").Contains("")
// true

str.New("golang").Contains("go ")
// false

str.New("lorem ipsum dolor").Contains("um do")
// true

```
