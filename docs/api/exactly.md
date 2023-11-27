# Exactly `bool`

`Exactly` returns true if the current string is exactly the same as the given string.

### Signature

```go
func (s *Str) Exactly(value string) bool
```

### Examples

```go
str.New("golang").Exactly("golang")
// true

str.New("golang").Exactly("")
// false

str.New("golang").Exactly("go")
// false

str.New("你好世界").Exactly("好")
// false

str.New("你好世界").Exactly("你好世界")
// true

```
