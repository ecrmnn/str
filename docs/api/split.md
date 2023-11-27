# Split `[]string`

`Split` returns a slice of strings containing the substrings of the current string separated by the given separator.

### Signature

```go
func (s *Str) Split(separator string) []string
```

### Examples

```go
str.New("go").Split("")
// []string{"g", "o"}

str.New("go go go").Split(" ")
// []string{"go", "go", "go"}

str.New("golang").Split(" ")
// []string{"golang"}

str.New("lorem ipsum").Split("_")
// []string{"lorem ipsum"}

str.New("").Split("")
// []string{}

str.New("    go  go  go").Split("")
// []string{" ", " ", " ", " ", "g", "o", " ", " ", "g", "o", " ", " ", "g", "o"}

str.New("✨🔥✨🔥✨").Split("🔥")
// []string{"✨", "✨", "✨"}

```
