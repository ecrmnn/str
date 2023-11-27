# ContainsAny `bool`

`ContainsAny` returns true if the current string contains any of the given substrings.

### Signature

```go
func (s *Str) ContainsAny(substring []string) bool
```

### Examples

```go
str.New("golang").
	ContainsAny([]string{""}),
// true

str.New("golang").
	ContainsAny([]string{}),
// false

str.New("lorem ipsum dolor").
	ContainsAny([]string{"lorem", "sit", "amet"}),
// true

str.New("lorem ipsum dolor").
	ContainsAny([]string{"sit", "amet"}),
// false

```
