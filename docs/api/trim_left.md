# TrimLeft `*Str`

`TrimLeft` returns a new instance of Str with all leading spaces removed.

### Signature

```go
func (s *Str) TrimLeft() *Str
```

### Examples

```go
str.New(" Hello,   World  ! ").TrimLeft().String()
// "Hello,   World  ! "

str.New("Hello,\tWorld ! ").TrimLeft().String()
// "Hello,\tWorld ! "

str.New(" \t\n\t Hello,\tWorld\n!\n\t").TrimLeft().String()
// "Hello,\tWorld\n!\n\t"

```
