# Capitalize `*Str`

`Capitalize` returns a new instance of Str with the first character capitalized.

### Signature

```go
func (s *Str) Capitalize() *Str
```

### Examples

```go
str.New("hello").Capitalize().String()
// "Hello"

str.New("_ _ _").Capitalize().String()
// "_ _ _"

str.New("hello World").Capitalize().String()
// "Hello World"

```
