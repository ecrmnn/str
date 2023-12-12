# Reverse `*Str`

`Reverse` returns a new instance of Str with the current string reversed.

### Signature

```go
func (s *Str) Reverse() *Str
```

### Examples

```go
str.New("i like c++").Reverse().String()
// "++c ekil i"

str.New("rocket 🚀 man").Reverse().String()
// "nam 🚀 tekcor"

str.New("go, go, go!").Reverse().String()
// "!og ,og ,og"

str.New("你好世界").Reverse().String()
// "界世好你"

str.New("你好世界").Reverse().String()
// "界世好你"

```
