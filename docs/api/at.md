# At `string`

`At` returns the character at the given index as a string.

### Signature

```go
func (s *Str) At(index int) string
```

### Examples

```go
str.New("golang").At(2)
// "l"

str.New("hello").At(0)
// "h"

str.New("æøå").At(1)
// "ø"

str.New("🚀x🌍").At(2)
// "🌍"

str.New("").At(0)
// ""

str.New("test").At(10)
// ""

```
