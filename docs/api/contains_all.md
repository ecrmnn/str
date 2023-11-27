# ContainsAll `bool`

`ContainsAll` returns true if the current string contains all of the given substrings.

### Signature

```go
func (s *Str) ContainsAll(substrings []string) bool
```

### Examples

```go
str.New("golang").
	ContainsAll([]string{"lan"}),
// true

str.New("golang").
	ContainsAll([]string{}),
// false

str.New("lorem ipsum dolor sit amet").
	ContainsAll([]string{"lorem", "sit", "amet"}),
// true

str.New("lorem ipsum dolor sit").
	ContainsAll([]string{"sit", "amet"}),
// false

```
