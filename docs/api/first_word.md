# FirstWord `*Str`

`FirstWord` returns a new instance of Str containing the first word of the current string.

### Signature

```go
func (s *Str) FirstWord() *Str
```

### Examples

```go
str.New("golang").FirstWord().String()
// "golang"

str.New("hello world").FirstWord().String()
// "hello"

str.New("    hello it    should   trim ").FirstWord().String()
// "hello"

str.New("æ ø å").FirstWord().String()
// "æ"

str.New("").FirstWord().String()
// ""

str.New("...hello").FirstWord().String()
// "...hello"

```
