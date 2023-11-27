# TrimRight `*Str`

`TrimRight` returns a new instance of Str with all trailing spaces removed.

### Signature

```go
func (s *Str) TrimRight() *Str
```

### Examples

```go
str.New(" Hello,   World  ! ").TrimRight().String()
// " Hello,   World  !"

str.New("Hello,\tWorld ! ").TrimRight().String()
// "Hello,\tWorld !"

str.New(" \t\n\t Hello,\tWorld\n!\n\t").TrimRight().String()
// " \t\n\t Hello,\tWorld\n!"

```
