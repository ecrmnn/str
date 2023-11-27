# Prepend `*Str`

`Prepend` prepends the given string to the beginning of the current string and returns a pointer to the current instance of Str.

### Signature

```go
func (s *Str) Prepend(prependStr string) *Str
```

### Examples

```go
str.New("♥️ is").Prepend("golang").String()
// "golang♥️ is"

str.New("!@#$%^&*()").Prepend("hello").String()
// "hello!@#$%^&*()"

str.New("").Prepend("øæå").String()
// "øæå"

str.New("🚀").Prepend("test").String()
// "test🚀"

```
