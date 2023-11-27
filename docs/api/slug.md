# Slug `*Str`

`Slug` returns a new Str with slugified version of the current string. The separator is used to replace all non-alphanumeric characters.

### Signature

```go
func (s *Str) Slug(separator string) *Str
```

### Examples

```go
str.New("  ❤️ go ❤️  ").Slug("_").String()
// "go"

str.New("  go ❤️  ").Slug("_").String()
// "go"

str.New("GoLang").Slug("_").String()
// "golang"

str.New("foo_@_bar-baz!").Slug("_").String()
// "foo_bar_baz"

str.New("foo_@_bar-baz!").Slug("-").String()
// "foo-bar-baz"

str.New("foo_@_bar-baz!").Slug("/").String()
// "foo/bar/baz"

str.New("").Slug("_").String()
// ""

str.New("goLang").Slug("_").String()
// "golang"

str.New("✨🔥✨🔥✨").Slug("_").String()
// ""

str.New("✨ 🔥 ✨ 🔥 ✨").Slug("_").String()
// ""

str.New("__hello__world__").Slug("_").String()
// "hello_world"

str.New("__hello__world__").Slug("-").String()
// "hello-world"

```
