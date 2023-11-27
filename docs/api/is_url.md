# IsUrl `bool`

`IsUrl` returns true if the current string is a valid URL.

### Signature

```go
func (s *Str) IsUrl() bool
```

### Examples

```go
str.New("http://example.com").IsUrl()
// true

str.New("https://subdomain.example.com").IsUrl()
// true

str.New("https://example.com/path").IsUrl()
// true

str.New("https://example.com/path?query=value").IsUrl()
// true

str.New("https://example.com/path#section").IsUrl()
// true

str.New("https://user:password@example.com").IsUrl()
// true

str.New("https://example.com:8080").IsUrl()
// true

str.New("https://my-domain-example.com").IsUrl()
// true

str.New("https://my_domain_example.com").IsUrl()
// false

str.New("https://192.168.0.1").IsUrl()
// true

str.New("ftp://example.com").IsUrl()
// true

str.New("example.com").IsUrl()
// false

str.New("https://").IsUrl()
// false

str.New("https://example.com/").IsUrl()
// true

str.New("https://example!.com").IsUrl()
// false

str.New("https:/example.com").IsUrl()
// false

str.New("https//example.com").IsUrl()
// false

str.New("https://example .com").IsUrl()
// false

str.New("https://localhost:8080").IsUrl()
// true

str.New("https://example&com").IsUrl()
// false

str.New("192.168.0.0").IsUrl()
// false

str.New("http://192.168.0.0").IsUrl()
// true

str.New("https://192.168.0").IsUrl()
// true

str.New("https://🔥.rs").IsUrl()
// true

```
