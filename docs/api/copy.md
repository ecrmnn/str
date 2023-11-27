# Copy `*Str`

`Copy` returns a new instance of Str with the same value as the current instance.

### Signature

```go
func (s *Str) Copy() *Str
```

### Examples

```go
func() bool {
	s := str.New("go")
	s2 := s.Copy().Append("lang")
	// s => go
	// s2 => golang
	return s.String() == s2.String()
}()
// false

```
