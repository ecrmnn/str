# String `string`

`String` returns the current value as a string.

### Signature

```go
func (s *Str) String() string
```

### Examples

```go
str.New("go").Capitalize().String()
// "Go"

str.New("").Lower().String()
// ""

```
