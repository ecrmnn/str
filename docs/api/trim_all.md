# TrimAll `*Str`

`TrimAll` returns a new instance of Str with all consecutive spaces replaced with a single space.

### Signature

```go
func (s *Str) TrimAll() *Str
```

### Examples

```go
str.New(" Hello,   World  ! ").TrimAll().String()
// "Hello, World !"

str.New("Hello,\tWorld ! ").TrimAll().String()
// "Hello, World !"

str.New(" \t\n\t Hello,\tWorld\n!\n\t").TrimAll().String()
// "Hello, World !"

```
