# IsEmpty `bool`

`IsEmpty` returns true if the current string is empty.

### Signature

```go
func (s *Str) IsEmpty() bool
```

### Examples

```go
str.New("golang").IsEmpty()
// false

str.New("").IsEmpty()
// true

str.New("   ").IsEmpty()
// false

```
