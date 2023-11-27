# Length `int`

`Length` returns the length of the current string.

### Signature

```go
func (s *Str) Length() int
```

### Examples

```go
str.New("go🔥").Length()
// 3

str.New("golang").Length()
// 6

str.New("lorem ipsum dolor").Length()
// 17

str.New("你好世界").Length()
// 4

str.New("你好界").Length()
// 3

```
