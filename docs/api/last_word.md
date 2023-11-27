# LastWord `*Str`

`LastWord` returns a new instance of Str containing the last word of the current string.

### Signature

```go
func (s *Str) LastWord() *Str
```

### Examples

```go
str.New("golang").LastWord().String()
// "golang"

str.New("hello world").LastWord().String()
// "world"

str.New("    hello it    should   trim ").LastWord().String()
// "trim"

str.New("æ ø å").LastWord().String()
// "å"

str.New("").LastWord().String()
// ""

str.New("...hello").LastWord().String()
// "...hello"

```
