# Append `*Str`

`Append` appends the given string to the end of the current string and returns a pointer to the current instance of Str.

### Signature

```go
func (s *Str) Append(value string) *Str
```

### Examples

```go
str.New("golang").Append(" is ♥️").String()
// "golang is ♥️"

str.New("hello").Append("!@#$%^&*()").String()
// "hello!@#$%^&*()"

str.New("").Append("øæå").String()
// "øæå"

str.New("test").Append("🚀").String()
// "test🚀"

```
