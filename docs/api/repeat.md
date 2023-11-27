# Repeat `*Str`

`Repeat` returns a new instance of Str with the current string repeated the given number of times.

### Signature

```go
func (s *Str) Repeat(count int) *Str
```

### Examples

```go
str.New("golang").Repeat(3).String()
// "golanggolanggolang"

str.New("golang").Repeat(0).String()
// ""

str.New("lorem ipsum").Repeat(2).String()
// "lorem ipsumlorem ipsum"

str.New("").Repeat(5).String()
// ""

```
