# TrimBoth `*Str`

`TrimBoth` returns a new instance of Str with all leading and trailing spaces removed.

### Signature

```go
func (s *Str) TrimBoth() *Str
```

### Examples

```go
str.New(" Hello,   World  ! ").TrimBoth().String()
// "Hello,   World  !"

str.New("Hello,\tWorld ! ").TrimBoth().String()
// "Hello,\tWorld !"

str.New(" \t\n\t Hello,\tWorld\n!\n\t").TrimBoth().String()
// "Hello,\tWorld\n!"

```
