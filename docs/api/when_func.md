# WhenFunc `*Str`

`WhenFunc` invokes the given callback if the given condition is true and returns a pointer to the current instance of Str. The condition function is passed a copy of Str.

### Signature

```go
func (s *Str) WhenFunc(condition func(s *Str) bool, callback func(s *Str) *Str) *Str
```

### Examples

```go
str.New("hello world").
	WhenFunc(func(s *str.Str) bool {
		return s.FirstWord().Exactly("hello")
	}, func(s *str.Str) *str.Str {
		return s.Append("!")
	}).
	String()
// "hello world!"

```
