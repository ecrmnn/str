# Camel `*Str`

`Camel` returns a new instance of Str in camelCase.

### Signature

```go
func (s *Str) Camel() *Str
```

### Examples

```go
str.New("string manipulation").Camel().String()
// "stringManipulation"

str.New("").Camel().String()
// ""

str.New("Hello world").Camel().String()
// "helloWorld"

```
