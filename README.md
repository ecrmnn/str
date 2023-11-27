![str](./docs/assets/str.png)

# str ![GitHub Workflow Status (with event)](https://img.shields.io/github/actions/workflow/status/ecrmnn/str/go.yml)

## A fluent string manipulation library in Go

### Installation

```bash
go get -u github.com/ecrmnn/str
```

### Documentation

Visit https://str.danieleckermann.com for full documentation

### API

- [After](#after-str)
- [AfterLast](#afterlast-str)
- [Append](#append-str)
- [At](#at-string)
- [Before](#before-str)
- [BeforeLast](#beforelast-str)
- [Between](#between-str)
- [Camel](#camel-str)
- [Capitalize](#capitalize-str)
- [Contains](#contains-bool)
- [ContainsAny](#containsany-bool)
- [ContainsAll](#containsall-bool)
- [Copy](#copy-str)
- [Count](#count-int)
- [EndsWith](#endswith-bool)
- [Exactly](#exactly-bool)
- [First](#first-str)
- [FirstWord](#firstword-str)
- [Last](#last-str)
- [LastWord](#lastword-str)
- [Length](#length-int)
- [Lower](#lower-str)
- [IsEmpty](#isempty-bool)
- [IsJson](#isjson-bool)
- [IsNotEmpty](#isnotempty-bool)
- [IsUrl](#isurl-bool)
- [Kebab](#kebab-str)
- [Pascal](#pascal-str)
- [Prepend](#prepend-str)
- [Repeat](#repeat-str)
- [Remove](#remove-str)
- [Replace](#replace-str)
- [Screaming](#screaming-str)
- [Slug](#slug-str)
- [Snake](#snake-str)
- [Split](#split-string)
- [StartsWith](#startswith-bool)
- [String](#string-string)
- [Swap](#swap-str)
- [Title](#title-str)
- [TrimAll](#trimall-str)
- [TrimBoth](#trimboth-str)
- [TrimLeft](#trimleft-str)
- [TrimRight](#trimright-str)
- [Upper](#upper-str)
- [When](#when-str)
- [WhenContains](#whencontains-str)
- [WhenContainsAny](#whencontainsany-str)
- [WhenContainsAll](#whencontainsall-str)
- [WhenEndsWith](#whenendswith-str)
- [WhenEmpty](#whenempty-str)
- [WhenExactly](#whenexactly-str)
- [WhenFunc](#whenfunc-str)
- [WhenNotEmpty](#whennotempty-str)
- [WhenNotExactly](#whennotexactly-str)
- [WhenStartsWith](#whenstartswith-str)
### After `*Str`

`After` returns a new instance of Str containing the substring after the first occurrence of the given substring.

```go
str.New("golang").After("go").String()
// "lang"

str.New("go go go!").After("go").String()
// " go go!"

str.New("øæå").After("æ").String()
// "å"

str.New("rocket 🚀🚀 man").After("🚀").String()
// "🚀 man"

```


### AfterLast `*Str`

`AfterLast` returns a new instance of Str containing the substring after the last occurrence of the given substring.

```go
str.New("golang").AfterLast("go").String()
// "lang"

str.New("go go go!").AfterLast("go").String()
// "!"

str.New("øæå").AfterLast("ø").String()
// "æå"

str.New("rocket 🚀🚀 man").AfterLast("🚀").String()
// " man"

```


### Append `*Str`

`Append` appends the given string to the end of the current string and returns a pointer to the current instance of Str.

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


### At `string`

`At` returns the character at the given index as a string.

```go
str.New("golang").At(2)
// "l"

str.New("hello").At(0)
// "h"

str.New("æøå").At(1)
// "ø"

str.New("🚀x🌍").At(2)
// "🌍"

str.New("").At(0)
// ""

str.New("test").At(10)
// ""

```


### Before `*Str`

`Before` returns a new instance of Str containing the substring before the first occurrence of the given substring.

```go
str.New("golang").Before("go").String()
// ""

str.New("go go go!").Before("go!").String()
// "go go "

str.New("øæå").Before("æ").String()
// "ø"

str.New("rocket 🚀🚀 man").Before("🚀").String()
// "rocket "

str.New("aaa").Before("a").String()
// ""

str.New("aabbcc").Before("bc").String()
// "aab"

```


### BeforeLast `*Str`

`BeforeLast` returns a new instance of Str containing the substring before the last occurrence of the given substring.

```go
str.New("golang").BeforeLast("go").String()
// ""

str.New("go go go!").BeforeLast("go").String()
// "go go "

str.New("øæøåøåø").BeforeLast("ø").String()
// "øæøåøå"

str.New("rocket 🚀🚀 man").BeforeLast("🚀").String()
// "rocket 🚀"

```


### Between `*Str`

`Between` returns a new instance of Str containing the substring between the first occurrence of the "from" substring and the first occurrence of the "to" substring.

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


### Camel `*Str`

`Camel` returns a new instance of Str in camelCase.

```go
str.New("string manipulation").Camel().String()
// "stringManipulation"

str.New("").Camel().String()
// ""

str.New("Hello world").Camel().String()
// "helloWorld"

```


### Capitalize `*Str`

`Capitalize` returns a new instance of Str with the first character capitalized.

```go
str.New("hello").Capitalize().String()
// "Hello"

str.New("_ _ _").Capitalize().String()
// "_ _ _"

str.New("hello World").Capitalize().String()
// "Hello World"

```


### Contains `bool`

`Contains` returns true if the current string contains the given substring.

```go
str.New("golang").Contains("lan")
// true

str.New("golang").Contains("")
// true

str.New("golang").Contains("go ")
// false

str.New("lorem ipsum dolor").Contains("um do")
// true

```


### ContainsAny `bool`

`ContainsAny` returns true if the current string contains any of the given substrings.

```go
str.New("golang").
	ContainsAny([]string{""}),
// true

str.New("golang").
	ContainsAny([]string{}),
// false

str.New("lorem ipsum dolor").
	ContainsAny([]string{"lorem", "sit", "amet"}),
// true

str.New("lorem ipsum dolor").
	ContainsAny([]string{"sit", "amet"}),
// false

```


### ContainsAll `bool`

`ContainsAll` returns true if the current string contains all of the given substrings.

```go
str.New("golang").
	ContainsAll([]string{"lan"}),
// true

str.New("golang").
	ContainsAll([]string{}),
// false

str.New("lorem ipsum dolor sit amet").
	ContainsAll([]string{"lorem", "sit", "amet"}),
// true

str.New("lorem ipsum dolor sit").
	ContainsAll([]string{"sit", "amet"}),
// false

```


### Copy `*Str`

`Copy` returns a new instance of Str with the same value as the current instance.

```go
func() bool {
	s := str.New("go")
	s2 := s.Copy().Append("lang")
	// s => go
	// s2 => golang
	return s.String() == s2.String()
}()
// false

```


### Count `int`

`Count` returns the number of occurrences of the given substring in the current string.

```go
str.New("golang").Count("lan")
// 1

str.New("golang").Count("")
// 7

str.New("golang").Count("go ")
// 0

str.New("lorem ipsum dolor").Count("um do")
// 1

str.New("你好世界").Count("好")
// 1

str.New("你好世界").Count("")
// 5

str.New("你好世界").Count("你好")
// 1

str.New("你好世界").Count("再见")
// 0

```


### EndsWith `bool`

`EndsWith` returns true if the current string ends with the given substring.

```go
str.New("golang").EndsWith("")
// false

str.New("golang").EndsWith("lang")
// true

str.New(" go ").EndsWith(" ")
// true

str.New("🔥🔥").EndsWith("🔥")
// true

str.New("✨🔥✨🔥✨").EndsWith("🔥")
// false

str.New("go").EndsWith("gox")
// false

str.New("").EndsWith("something")
// false

str.New("abc").EndsWith("")
// false

str.New("abc").EndsWith("abcdef")
// false

str.New("Golang").EndsWith("lang")
// true

str.New("Golang").EndsWith("LANG")
// false

str.New("café").EndsWith("fé")
// true

str.New("abcabc").EndsWith("abc")
// true

str.New("abcabc").EndsWith("bc")
// true

```


### Exactly `bool`

`Exactly` returns true if the current string is exactly the same as the given string.

```go
str.New("golang").Exactly("golang")
// true

str.New("golang").Exactly("")
// false

str.New("golang").Exactly("go")
// false

str.New("你好世界").Exactly("好")
// false

str.New("你好世界").Exactly("你好世界")
// true

```


### First `*Str`

`First` returns s new instance of Str containing the first character of the current string.

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


### FirstWord `*Str`

`FirstWord` returns a new instance of Str containing the first word of the current string.

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


### Last `*Str`

`Last` returns s new instance of Str containing the last character of the current string.

```go
str.New("golang").Last().String()
// "g"

str.New("hello world").Last().String()
// "d"

str.New("🔥 Go 🤩✨").Last().String()
// "✨"

str.New("æøå").Last().String()
// "å"

str.New("").Last().String()
// ""

```


### LastWord `*Str`

`LastWord` returns a new instance of Str containing the last word of the current string.

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


### Length `int`

`Length` returns the length of the current string.

```go
str.New("go🔥").Length()
// 3

str.New("golang").Length()
// 6

str.New("lorem ipsum dolor").Length()
// 17

str.New("你好世界").Length()
// 4

str.New("你好界").Length()
// 3

```


### Lower `*Str`

`Lower` returns a new instance of Str with all characters in lowercase.

```go
str.New("").Lower().String()
// ""

str.New("✨").Lower().String()
// "✨"

str.New("&").Lower().String()
// "&"

str.New("Æ").Lower().String()
// "æ"

str.New("TEST").Lower().String()
// "test"

```


### IsEmpty `bool`

`IsEmpty` returns true if the current string is empty.

```go
str.New("golang").IsEmpty()
// false

str.New("").IsEmpty()
// true

str.New("   ").IsEmpty()
// false

```


### IsJson `bool`

`IsJson` returns true if the current string is valid JSON.

```go
str.New(`{"key": "value"}`).IsJson()
// true

str.New(`{"key": "value", "nested": {"inner": 42}}`).IsJson()
// true

str.New("not a valid JSON").IsJson()
// false

str.New("").IsJson()
// false

```


### IsNotEmpty `bool`

`IsNotEmpty` returns true if the current string is not empty.

```go
str.New("golang").IsNotEmpty()
// true

str.New("").IsNotEmpty()
// false

str.New("   ").IsNotEmpty()
// true

```


### IsUrl `bool`

`IsUrl` returns true if the current string is a valid URL.

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


### Kebab `*Str`

`Kebab` returns a new lowercased instance of Str with all spaces and underscores replaced with hyphens.

```go
str.New("").Kebab().String()
// ""

str.New("go go go").Kebab().String()
// "go-go-go"

str.New("GoLang").Kebab().String()
// "go-lang"

str.New("goLang").Kebab().String()
// "go-lang"

str.New("✨🔥✨🔥✨").Kebab().String()
// "✨🔥✨🔥✨"

str.New("✨ 🔥 ✨ 🔥 ✨").Kebab().String()
// "✨-🔥-✨-🔥-✨"

```


### Pascal `*Str`

`Pascal` returns a new instance of Str in PascalCase.

```go
str.New("").Pascal().String()
// ""

str.New("go go go").Pascal().String()
// "GoGoGo"

str.New("hello world").Pascal().String()
// "HelloWorld"

str.New("Go-lang").Pascal().String()
// "Go-lang"

str.New("goLang").Pascal().String()
// "GoLang"

str.New("✨🔥✨🔥✨").Pascal().String()
// "✨🔥✨🔥✨"

str.New("✨ 🔥 ✨ 🔥 ✨").Pascal().String()
// "✨🔥✨🔥✨"

```


### Prepend `*Str`

`Prepend` prepends the given string to the beginning of the current string and returns a pointer to the current instance of Str.

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


### Repeat `*Str`

`Repeat` returns a new instance of Str with the current string repeated the given number of times.

```go
str.New("golang").Repeat(3).String()
// "golanggolanggolang"

str.New("golang").Repeat(0).String()
// ""

str.New("lorem ipsum").Repeat(2).String()
// "lorem ipsumlorem ipsum"

str.New("").Repeat(5).String()
// ""

```


### Remove `*Str`

`Remove` returns a new instance of Str with all occurrences of the given substring removed.

```go
str.New("i like c++").Remove("c++").String()
// "i like "

str.New("rocket 🚀 man").Remove("🚀").String()
// "rocket  man"

str.New("go, go, go!").Remove("go").String()
// ", , !"

str.New("你好世界").Remove("你好").String()
// "世界"

str.New("").Remove("go").String()
// ""

```


### Replace `*Str`

`Replace` returns a new instance of Str with all occurrences of the given substring replaced with the given replacement string.

```go
str.New("i like c++").Replace("c++", "go").String()
// "i like go"

str.New("rocket 🚀 man").Replace("🚀", "🌎").String()
// "rocket 🌎 man"

str.New("go, go, go!").Replace("go", "Go").String()
// "Go, Go, Go!"

str.New("你好世界").Replace("你好", "再见").String()
// "再见世界"

str.New("你好世界").Replace("你好", "").String()
// "世界"

```


### Screaming `*Str`



```go
str.New("go go go").Screaming().String()
// "GO_GO_GO"

str.New("GoLang").Screaming().String()
// "GO_LANG"

str.New("ALREADY_SCREAMING").Screaming().String()
// "ALREADY_SCREAMING"

str.New("").Screaming().String()
// ""

str.New("goLang").Screaming().String()
// "GO_LANG"

str.New("✨🔥✨🔥✨").Screaming().String()
// "✨🔥✨🔥✨"

str.New("✨ 🔥 ✨ 🔥 ✨").Screaming().String()
// "✨_🔥_✨_🔥_✨"

```


### Slug `*Str`

`Slug` returns a new Str with slugified version of the current string. The separator is used to replace all non-alphanumeric characters.

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


### Snake `*Str`

`Snake` returns a new lowercased instance of Str with all spaces and hyphens replaced with underscores.

```go
str.New("go go go").Snake().String()
// "go_go_go"

str.New("GoLang").Snake().String()
// "go_lang"

str.New("already_snake_cased").Snake().String()
// "already_snake_cased"

str.New("").Snake().String()
// ""

str.New("goLang").Snake().String()
// "go_lang"

str.New("✨🔥✨🔥✨").Snake().String()
// "✨🔥✨🔥✨"

str.New("✨ 🔥 ✨ 🔥 ✨").Snake().String()
// "✨_🔥_✨_🔥_✨"

```


### Split `[]string`

`Split` returns a slice of strings containing the substrings of the current string separated by the given separator.

```go
str.New("go").Split("")
// []string{"g", "o"}

str.New("go go go").Split(" ")
// []string{"go", "go", "go"}

str.New("golang").Split(" ")
// []string{"golang"}

str.New("lorem ipsum").Split("_")
// []string{"lorem ipsum"}

str.New("").Split("")
// []string{}

str.New("    go  go  go").Split("")
// []string{" ", " ", " ", " ", "g", "o", " ", " ", "g", "o", " ", " ", "g", "o"}

str.New("✨🔥✨🔥✨").Split("🔥")
// []string{"✨", "✨", "✨"}

```


### StartsWith `bool`

`StartsWith` returns true if the current string starts with the given substring.

```go
str.New("golang").StartsWith("")
// false

str.New("golang").StartsWith("go")
// true

str.New(" go ").StartsWith(" ")
// true

str.New("🔥🔥").StartsWith("🔥")
// true

str.New("✨🔥✨🔥✨").StartsWith("🔥")
// false

```


### String `string`

`String` returns the current value as a string.

```go
str.New("go").Capitalize().String()
// "Go"

str.New("").Lower().String()
// ""

```


### Swap `*Str`

`Swap` returns a new instance of Str where all occurrences of keys are replaces with their corresponding values.

```go
str.New("I love Rust").
	Swap(
		map[string]string{
			"Rust": "Go",
		}).
	String()
// "I love Go"

str.New("Rust. Rust! RUST!!").
	Swap(
		map[string]string{
			"Rust": "Go",
		}).
	String()
// "Go. Go! RUST!!"

str.New("Rust. Rust! RUST!!").
	Swap(
		map[string]string{
			"Rust": "Go",
			"RUST": "Go",
		}).
	String()
// "Go. Go! Go!!"

str.New("I ❤️ U").
	Swap(
		map[string]string{
			"❤️": "🤬",
			"I":  "You",
			"U":  "Me",
		}).
	String()
// "You 🤬 Me"

str.New("BLÅBÆRSYLTETØY blåbærsyltetøy").
	Swap(
		map[string]string{
			"Æ": "",
			"Ø": "",
			"Å": "",
			"æ": "",
			"ø": "",
			"å": "",
		}).
	String()
// "BLBRSYLTETY blbrsyltety"

str.New("").
	Swap(
		map[string]string{
			"Æ": "",
			"Ø": "",
			"Å": "",
			"æ": "",
			"ø": "",
			"å": "",
		}).
	String()
// ""

```


### Title `*Str`

`Title` returns a new instance of Str with the first letter of each word capitalized.

```go
str.New("go øl åre ære").Title().String()
// "Go Øl Åre Ære"

str.New("go ✨ go").Title().String()
// "Go ✨ Go"

str.New("_ _ _").Title().String()
// "_ _ _"

str.New("").Title().String()
// ""

```


### TrimAll `*Str`

`TrimAll` returns a new instance of Str with all consecutive spaces replaced with a single space.

```go
str.New(" Hello,   World  ! ").TrimAll().String()
// "Hello, World !"

str.New("Hello,\tWorld ! ").TrimAll().String()
// "Hello, World !"

str.New(" \t\n\t Hello,\tWorld\n!\n\t").TrimAll().String()
// "Hello, World !"

```


### TrimBoth `*Str`

`TrimBoth` returns a new instance of Str with all leading and trailing spaces removed.

```go
str.New(" Hello,   World  ! ").TrimBoth().String()
// "Hello,   World  !"

str.New("Hello,\tWorld ! ").TrimBoth().String()
// "Hello,\tWorld !"

str.New(" \t\n\t Hello,\tWorld\n!\n\t").TrimBoth().String()
// "Hello,\tWorld\n!"

```


### TrimLeft `*Str`

`TrimLeft` returns a new instance of Str with all leading spaces removed.

```go
str.New(" Hello,   World  ! ").TrimLeft().String()
// "Hello,   World  ! "

str.New("Hello,\tWorld ! ").TrimLeft().String()
// "Hello,\tWorld ! "

str.New(" \t\n\t Hello,\tWorld\n!\n\t").TrimLeft().String()
// "Hello,\tWorld\n!\n\t"

```


### TrimRight `*Str`

`TrimRight` returns a new instance of Str with all trailing spaces removed.

```go
str.New(" Hello,   World  ! ").TrimRight().String()
// " Hello,   World  !"

str.New("Hello,\tWorld ! ").TrimRight().String()
// "Hello,\tWorld !"

str.New(" \t\n\t Hello,\tWorld\n!\n\t").TrimRight().String()
// " \t\n\t Hello,\tWorld\n!"

```


### Upper `*Str`

`Upper` returns a new instance of Str with all characters in uppercase.

```go
str.New("").Upper().String()
// ""

str.New("✨").Upper().String()
// "✨"

str.New("&").Upper().String()
// "&"

str.New("æ").Upper().String()
// "Æ"

str.New("test").Upper().String()
// "TEST"

```


### When `*Str`

`When` invokes the given callback if the given condition is true and returns a pointer to the current instance of Str.

```go
str.New("go").
	When(true, func(s *str.Str) *str.Str {
		return s.Append("lang")
	}).
	String()
// "golang"

str.New("go").
	When(false, func(s *str.Str) *str.Str {
		return s.Append("lang")
	}).
	String()
// "go"

str.New("").
	When(false, func(s *str.Str) *str.Str {
		return s.Append("lang")
	}).
	String()
// ""

```


### WhenContains `*Str`

`WhenContains` invokes the given callback if the current string contains the given value and returns a pointer to the current instance of Str.

```go
str.New("i like go").
	WhenContains("go", func(s *str.Str) *str.Str {
		return s.Append("lang")
	}).
	String()
// "i like golang"

str.New("i like go").
	WhenContains("gox", func(s *str.Str) *str.Str {
		return s.Append("lang")
	}).
	String()
// "i like go"

str.New("").
	WhenContains("go", func(s *str.Str) *str.Str {
		return s.Append("lang")
	}).
	String()
// ""

```


### WhenContainsAny `*Str`

`WhenContainsAny` invokes the given callback if the current string contains any of the given values and returns a pointer to the current instance of Str.

```go
str.New("i like go").
	WhenContainsAny([]string{"go", "c++"}, func(s *str.Str) *str.Str {
		return s.Append("lang")
	}).
	String()
// "i like golang"

str.New("i like go").
	WhenContainsAny([]string{"a", "b", "c"}, func(s *str.Str) *str.Str {
		return s.Append("lang")
	}).
	String()
// "i like go"

str.New("").
	WhenContainsAny([]string{"go", "lang"}, func(s *str.Str) *str.Str {
		return s.Append("lang")
	}).
	String()
// ""

```


### WhenContainsAll `*Str`

`WhenContainsAll` invokes the given callback if the current string contains all of the given values and returns a pointer to the current instance of Str.

```go
str.New("i like go not c++").
	WhenContainsAll([]string{"go", "c++"}, func(s *str.Str) *str.Str {
		return s.Append("!")
	}).
	String()
// "i like go not c++!"

str.New("i like go").
	WhenContainsAll([]string{"a", "b", "c"}, func(s *str.Str) *str.Str {
		return s.Append("lang")
	}).
	String()
// "i like go"

str.New("").
	WhenContainsAll([]string{"go", "lang"}, func(s *str.Str) *str.Str {
		return s.Append("lang")
	}).
	String()
// ""

```


### WhenEndsWith `*Str`

`WhenEndsWith` invokes the given callback if the current string ends with the given value and returns a pointer to the current instance of Str.

```go
str.New("i like go").
	WhenEndsWith("go", func(s *str.Str) *str.Str {
		return s.Append("lang")
	}).
	String()
// "i like golang"

str.New("i like go").
	WhenEndsWith("gox", func(s *str.Str) *str.Str {
		return s.Append("lang")
	}).
	String()
// "i like go"

str.New("").
	WhenEndsWith("go", func(s *str.Str) *str.Str {
		return s.Append("lang")
	}).
	String()
// ""

```


### WhenEmpty `*Str`

`WhenEmpty` invokes the given closure if the current string is empty and returns a pointer to the current instance of Str.

```go
str.New("go").
	WhenEmpty(func(s *str.Str) *str.Str {
		return s.Append("lang")
	}).
	String()
// "go"

str.New("").
	WhenEmpty(func(s *str.Str) *str.Str {
		return s.Append("golang")
	}).
	String()
// "golang"

```


### WhenExactly `*Str`

`WhenExactly` invokes the given callback if the current string is exactly the same as the given value and returns a pointer to the current instance of Str.

```go
str.New("go").
	WhenExactly("go", func(s *str.Str) *str.Str {
		return s.Append("lang")
	}).
	String()
// "golang"

str.New("").
	WhenExactly("go", func(s *str.Str) *str.Str {
		return s.Append("golang")
	}).
	String()
// ""

str.New("🚀").
	WhenExactly("🚀", func(s *str.Str) *str.Str {
		return s.Append("✨")
	}).
	String()
// "🚀✨"

```


### WhenFunc `*Str`

`WhenFunc` invokes the given callback if the given condition is true and returns a pointer to the current instance of Str. The condition function is passed a copy of Str.

```go
str.New("hello world").
	WhenFunc(func(s *str.Str) bool {
		return s.FirstWord().Exactly("hello")
	}, func(s *str.Str) *str.Str {
		return s.Append("!")
	}).
	String()
// "hello world!"

```


### WhenNotEmpty `*Str`

`WhenNotEmpty` invokes the given closure if the current string is not empty and returns a pointer to the current instance of Str.

```go
str.New("go").
	WhenNotEmpty(func(s *str.Str) *str.Str {
		return s.Append("lang")
	}).
	String()
// "golang"

str.New("").
	WhenNotEmpty(func(s *str.Str) *str.Str {
		return s.Append("golang")
	}).
	String()
// ""

```


### WhenNotExactly `*Str`

`WhenNotExactly` invokes the given callback if the current string is not exactly the same as the given value and returns a pointer to the current instance of Str.

```go
str.New("go").
	WhenNotExactly("go", func(s *str.Str) *str.Str {
		return s.Append("lang")
	}).
	String()
// "go"

str.New("").
	WhenNotExactly("go", func(s *str.Str) *str.Str {
		return s.Append("golang")
	}).
	String()
// "golang"

str.New("🚀").
	WhenNotExactly("🚀", func(s *str.Str) *str.Str {
		return s.Append("✨")
	}).
	String()
// "🚀"

```


### WhenStartsWith `*Str`

`WhenStartsWith` invokes the given callback if the current string starts with the given value and returns a pointer to the current instance of Str.

```go
str.New("go").
	WhenStartsWith("go", func(s *str.Str) *str.Str {
		return s.Append("lang")
	}).
	String()
// "golang"

str.New("go").
	WhenStartsWith("gox", func(s *str.Str) *str.Str {
		return s.Append("lang")
	}).
	String()
// "go"

str.New("").
	WhenStartsWith("go", func(s *str.Str) *str.Str {
		return s.Append("lang")
	}).
	String()
// ""

```
