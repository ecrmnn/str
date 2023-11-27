# IsJson `bool`

`IsJson` returns true if the current string is valid JSON.

### Signature

```go
func (s *Str) IsJson() bool
```

### Examples

```go
str.New(`{"key": "value"}`).IsJson()
// true

str.New(`{"key": "value", "nested": {"inner": 42}}`).IsJson()
// true

str.New("not a valid JSON").IsJson()
// false

str.New("").IsJson()
// false

```
