# Count `int`

`Count` returns the number of occurrences of the given substring in the current string.

### Signature

```go
func (s *Str) Count(substring string) int
```

### Examples

```go
str.New("golang").Count("lan")
// 1

str.New("golang").Count("")
// 7

str.New("golang").Count("go ")
// 0

str.New("lorem ipsum dolor").Count("um do")
// 1

str.New("你好世界").Count("好")
// 1

str.New("你好世界").Count("")
// 5

str.New("你好世界").Count("你好")
// 1

str.New("你好世界").Count("再见")
// 0

```
