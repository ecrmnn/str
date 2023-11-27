# EndsWith `bool`

`EndsWith` returns true if the current string ends with the given substring.

### Signature

```go
func (s *Str) EndsWith(substring string) bool
```

### Examples

```go
str.New("golang").EndsWith("")
// false

str.New("golang").EndsWith("lang")
// true

str.New(" go ").EndsWith(" ")
// true

str.New("🔥🔥").EndsWith("🔥")
// true

str.New("✨🔥✨🔥✨").EndsWith("🔥")
// false

str.New("go").EndsWith("gox")
// false

str.New("").EndsWith("something")
// false

str.New("abc").EndsWith("")
// false

str.New("abc").EndsWith("abcdef")
// false

str.New("Golang").EndsWith("lang")
// true

str.New("Golang").EndsWith("LANG")
// false

str.New("café").EndsWith("fé")
// true

str.New("abcabc").EndsWith("abc")
// true

str.New("abcabc").EndsWith("bc")
// true

```
