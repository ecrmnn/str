# Between `*Str`

`Between` returns a new instance of Str containing the substring between the first occurrence of the "from" substring and the first occurrence of the "to" substring.

### Signature

```go
func (s *Str) Between(from string, to string) *Str
```

### Examples

```go
str.New("{{hello}}").Between("{{", "}}").String()
// "hello"

str.New("🚀✨🔥").Between("🚀", "🔥").String()
// "✨"

str.New("!#$%").Between("$", "%").String()
// ""

str.New("rocket 🚀🚀 man").Between("rocket", "man").String()
// " 🚀🚀 "

str.New("abc").Between("", "c").String()
// "ab"

str.New("abc").Between("a", "").String()
// ""

str.New("abc").Between("", "").String()
// ""

str.New("abc").Between("a", "c").String()
// "b"

str.New("dddabc").Between("a", "c").String()
// "b"

str.New("abcddd").Between("a", "c").String()
// "b"

str.New("dddabcddd").Between("a", "c").String()
// "b"

str.New("hannah").Between("ha", "ah").String()
// "nn"

str.New("[a]ab[b]").Between("[", "]").String()
// "a"

str.New("foofoobar").Between("foo", "bar").String()
// "foo"

str.New("foobarbar").Between("foo", "bar").String()
// ""

```
