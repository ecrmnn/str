# IsNotEmpty `bool`

`IsNotEmpty` returns true if the current string is not empty.

### Signature

```go
func (s *Str) IsNotEmpty() bool
```

### Examples

```go
str.New("golang").IsNotEmpty()
// true

str.New("").IsNotEmpty()
// false

str.New("   ").IsNotEmpty()
// true

```
