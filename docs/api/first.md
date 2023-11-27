# First `*Str`

`First` returns s new instance of Str containing the first character of the current string.

### Signature

```go
func (s *Str) First() *Str
```

### Examples

```go
str.New("golang").First().String()
// "g"

str.New("hello world").First().String()
// "h"

str.New("æ ø å").First().String()
// "æ"

str.New("😅👋").First().String()
// "😅"

str.New("").First().String()
// ""

str.New("こんにちは").First().String()
// "こ"

str.New("你好").First().String()
// "你"

```
