# StartsWith `bool`

`StartsWith` returns true if the current string starts with the given substring.

### Signature

```go
func (s *Str) StartsWith(substring string) bool
```

### Examples

```go
str.New("golang").StartsWith("")
// false

str.New("golang").StartsWith("go")
// true

str.New(" go ").StartsWith(" ")
// true

str.New("🔥🔥").StartsWith("🔥")
// true

str.New("✨🔥✨🔥✨").StartsWith("🔥")
// false

```
