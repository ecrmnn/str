# Clone `*Str`

`Clone` returns a new instance of Str with the same value as the current instance.

### Signature

```go
func (s *Str) Clone() *Str
```

### Examples

```go
func() bool {
	s := str.New("i like go")
	if s.Clone().LastWord().Exactly("go") {
		s.Append(" ❤️")
	}
	return s.String() == "i like go ❤️"
}()
// true

func() bool {
	s := str.New("go")
	s2 := s.Clone().Append("lang")
	// s => go
	// s2 => golang
	return s.String() == s2.String()
}()
// false

```
