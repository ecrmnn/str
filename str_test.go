package str_test

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"reflect"
	"strings"
	"testing"

	"github.com/ecrmnn/str"
)

type fixture struct {
	actual   any
	expected any
}

type testCase struct {
	name     string
	fixtures []fixture
}

var testCases = []testCase{
	{
		name: "After",
		fixtures: []fixture{
			{
				actual:   str.New("golang").After("go").String(),
				expected: "lang",
			},
			{
				actual:   str.New("go go go!").After("go").String(),
				expected: " go go!",
			},
			{
				actual:   str.New("øæå").After("æ").String(),
				expected: "å",
			},
			{
				actual:   str.New("rocket 🚀🚀 man").After("🚀").String(),
				expected: "🚀 man",
			},
		},
	},
	{
		name: "AfterLast",
		fixtures: []fixture{
			{
				actual:   str.New("golang").AfterLast("go").String(),
				expected: "lang",
			},
			{
				actual:   str.New("go go go!").AfterLast("go").String(),
				expected: "!",
			},
			{
				actual:   str.New("øæå").AfterLast("ø").String(),
				expected: "æå",
			},
			{
				actual:   str.New("rocket 🚀🚀 man").AfterLast("🚀").String(),
				expected: " man",
			},
		},
	},
	{
		name: "Append",
		fixtures: []fixture{
			{
				actual:   str.New("golang").Append(" is ♥️").String(),
				expected: "golang is ♥️",
			},
			{
				actual:   str.New("hello").Append("!@#$%^&*()").String(),
				expected: "hello!@#$%^&*()",
			},
			{
				actual:   str.New("").Append("øæå").String(),
				expected: "øæå",
			},
			{
				actual:   str.New("test").Append("🚀").String(),
				expected: "test🚀",
			},
		},
	},
	{
		name: "At",
		fixtures: []fixture{
			{
				actual:   str.New("golang").At(2),
				expected: "l",
			},
			{
				actual:   str.New("hello").At(0),
				expected: "h",
			},
			{
				actual:   str.New("æøå").At(1),
				expected: "ø",
			},
			{
				actual:   str.New("🚀x🌍").At(2),
				expected: "🌍",
			},
			{
				actual:   str.New("").At(0),
				expected: "",
			},
			{
				actual:   str.New("test").At(10),
				expected: "",
			},
		},
	},
	{
		name: "Before",
		fixtures: []fixture{
			{
				actual:   str.New("golang").Before("go").String(),
				expected: "",
			},
			{
				actual:   str.New("go go go!").Before("go!").String(),
				expected: "go go ",
			},
			{
				actual:   str.New("øæå").Before("æ").String(),
				expected: "ø",
			},
			{
				actual:   str.New("rocket 🚀🚀 man").Before("🚀").String(),
				expected: "rocket ",
			},
			{
				actual:   str.New("aaa").Before("a").String(),
				expected: "",
			},
			{
				actual:   str.New("aabbcc").Before("bc").String(),
				expected: "aab",
			},
		},
	},
	{
		name: "BeforeLast",
		fixtures: []fixture{
			{
				actual:   str.New("golang").BeforeLast("go").String(),
				expected: "",
			},
			{
				actual:   str.New("go go go!").BeforeLast("go").String(),
				expected: "go go ",
			},
			{
				actual:   str.New("øæøåøåø").BeforeLast("ø").String(),
				expected: "øæøåøå",
			},
			{
				actual:   str.New("rocket 🚀🚀 man").BeforeLast("🚀").String(),
				expected: "rocket 🚀",
			},
		},
	},
	{
		name: "Between",
		fixtures: []fixture{
			{
				actual:   str.New("{{hello}}").Between("{{", "}}").String(),
				expected: "hello",
			},
			{
				actual:   str.New("🚀✨🔥").Between("🚀", "🔥").String(),
				expected: "✨",
			},
			{
				actual:   str.New("!#$%").Between("$", "%").String(),
				expected: "",
			},
			{
				actual:   str.New("rocket 🚀🚀 man").Between("rocket", "man").String(),
				expected: " 🚀🚀 ",
			},
			{
				actual:   str.New("abc").Between("", "c").String(),
				expected: "ab",
			},
			{
				actual:   str.New("abc").Between("a", "").String(),
				expected: "",
			},
			{
				actual:   str.New("abc").Between("", "").String(),
				expected: "",
			},
			{
				actual:   str.New("abc").Between("a", "c").String(),
				expected: "b",
			},
			{
				actual:   str.New("dddabc").Between("a", "c").String(),
				expected: "b",
			},
			{
				actual:   str.New("abcddd").Between("a", "c").String(),
				expected: "b",
			},
			{
				actual:   str.New("dddabcddd").Between("a", "c").String(),
				expected: "b",
			},
			{
				actual:   str.New("hannah").Between("ha", "ah").String(),
				expected: "nn",
			},
			{
				actual:   str.New("[a]ab[b]").Between("[", "]").String(),
				expected: "a",
			},
			{
				actual:   str.New("foofoobar").Between("foo", "bar").String(),
				expected: "foo",
			},
			{
				actual:   str.New("foobarbar").Between("foo", "bar").String(),
				expected: "",
			},
		},
	},
	{
		name: "Camel",
		fixtures: []fixture{
			{
				actual:   str.New("string manipulation").Camel().String(),
				expected: "stringManipulation",
			},
			{
				actual:   str.New("").Camel().String(),
				expected: "",
			},
			{
				actual:   str.New("Hello world").Camel().String(),
				expected: "helloWorld",
			},
		},
	},
	{
		name: "Capitalize",
		fixtures: []fixture{
			{
				actual:   str.New("hello").Capitalize().String(),
				expected: "Hello",
			},
			{
				actual:   str.New("_ _ _").Capitalize().String(),
				expected: "_ _ _",
			},
			{
				actual:   str.New("hello World").Capitalize().String(),
				expected: "Hello World",
			},
		},
	},
	{
		name: "Contains",
		fixtures: []fixture{
			{
				actual:   str.New("golang").Contains("lan"),
				expected: true,
			},
			{
				actual:   str.New("golang").Contains(""),
				expected: true,
			},
			{
				actual:   str.New("golang").Contains("go "),
				expected: false,
			},
			{
				actual:   str.New("lorem ipsum dolor").Contains("um do"),
				expected: true,
			},
		},
	},
	{
		name: "ContainsAny",
		fixtures: []fixture{
			{
				actual: str.New("golang").
					ContainsAny([]string{""}),
				expected: true,
			},
			{
				actual: str.New("golang").
					ContainsAny([]string{}),
				expected: false,
			},
			{
				actual: str.New("lorem ipsum dolor").
					ContainsAny([]string{"lorem", "sit", "amet"}),
				expected: true,
			},
			{
				actual: str.New("lorem ipsum dolor").
					ContainsAny([]string{"sit", "amet"}),
				expected: false,
			},
		},
	},
	{
		name: "ContainsAll",
		fixtures: []fixture{
			{
				actual: str.New("golang").
					ContainsAll([]string{"lan"}),
				expected: true,
			},
			{
				actual: str.New("golang").
					ContainsAll([]string{}),
				expected: false,
			},
			{
				actual: str.New("lorem ipsum dolor sit amet").
					ContainsAll([]string{"lorem", "sit", "amet"}),
				expected: true,
			},
			{
				actual: str.New("lorem ipsum dolor sit").
					ContainsAll([]string{"sit", "amet"}),
				expected: false,
			},
		},
	},
	{
		name: "Clone",
		fixtures: []fixture{
			{
				actual: func() bool {
					s := str.New("i like go")
					if s.Clone().LastWord().Exactly("go") {
						s.Append(" ❤️")
					}
					return s.String() == "i like go ❤️"
				}(),
				expected: true,
			},
			{
				actual: func() bool {
					s := str.New("go")
					s2 := s.Clone().Append("lang")
					// s => go
					// s2 => golang
					return s.String() == s2.String()
				}(),
				expected: false,
			},
		},
	},
	{
		name: "Count",
		fixtures: []fixture{
			{
				actual:   str.New("golang").Count("lan"),
				expected: 1,
			},
			{
				actual:   str.New("golang").Count(""),
				expected: 7,
			},
			{
				actual:   str.New("golang").Count("go "),
				expected: 0,
			},
			{
				actual:   str.New("lorem ipsum dolor").Count("um do"),
				expected: 1,
			},
			{
				actual:   str.New("你好世界").Count("好"),
				expected: 1,
			},
			{
				actual:   str.New("你好世界").Count(""),
				expected: 5,
			},
			{
				actual:   str.New("你好世界").Count("你好"),
				expected: 1,
			},
			{
				actual:   str.New("你好世界").Count("再见"),
				expected: 0,
			},
		},
	},
	{
		name: "EndsWith",
		fixtures: []fixture{
			{
				actual:   str.New("golang").EndsWith(""),
				expected: false,
			},
			{
				actual:   str.New("golang").EndsWith("lang"),
				expected: true,
			},
			{
				actual:   str.New(" go ").EndsWith(" "),
				expected: true,
			},
			{
				actual:   str.New("🔥🔥").EndsWith("🔥"),
				expected: true,
			},
			{
				actual:   str.New("✨🔥✨🔥✨").EndsWith("🔥"),
				expected: false,
			},
			{
				actual:   str.New("go").EndsWith("gox"),
				expected: false,
			},
			{
				actual:   str.New("").EndsWith("something"),
				expected: false,
			},
			{
				actual:   str.New("abc").EndsWith(""),
				expected: false,
			},
			{
				actual:   str.New("abc").EndsWith("abcdef"),
				expected: false,
			},
			{
				actual:   str.New("Golang").EndsWith("lang"),
				expected: true,
			},
			{
				actual:   str.New("Golang").EndsWith("LANG"),
				expected: false,
			},
			{
				actual:   str.New("café").EndsWith("fé"),
				expected: true,
			},
			{
				actual:   str.New("abcabc").EndsWith("abc"),
				expected: true,
			},
			{
				actual:   str.New("abcabc").EndsWith("bc"),
				expected: true,
			},
		},
	},
	{
		name: "Exactly",
		fixtures: []fixture{
			{
				actual:   str.New("golang").Exactly("golang"),
				expected: true,
			},
			{
				actual:   str.New("golang").Exactly(""),
				expected: false,
			},
			{
				actual:   str.New("golang").Exactly("go"),
				expected: false,
			},
			{
				actual:   str.New("你好世界").Exactly("好"),
				expected: false,
			},
			{
				actual:   str.New("你好世界").Exactly("你好世界"),
				expected: true,
			},
		},
	},
	{
		name: "First",
		fixtures: []fixture{
			{
				actual:   str.New("golang").First().String(),
				expected: "g",
			},
			{
				actual:   str.New("hello world").First().String(),
				expected: "h",
			},
			{
				actual:   str.New("æ ø å").First().String(),
				expected: "æ",
			},
			{
				actual:   str.New("😅👋").First().String(),
				expected: "😅",
			},
			{
				actual:   str.New("").First().String(),
				expected: "",
			},
			{
				actual:   str.New("こんにちは").First().String(),
				expected: "こ",
			},
			{
				actual:   str.New("你好").First().String(),
				expected: "你",
			},
		},
	},
	{
		name: "FirstWord",
		fixtures: []fixture{
			{
				actual:   str.New("golang").FirstWord().String(),
				expected: "golang",
			},
			{
				actual:   str.New("hello world").FirstWord().String(),
				expected: "hello",
			},
			{
				actual:   str.New("    hello it    should   trim ").FirstWord().String(),
				expected: "hello",
			},
			{
				actual:   str.New("æ ø å").FirstWord().String(),
				expected: "æ",
			},
			{
				actual:   str.New("").FirstWord().String(),
				expected: "",
			},
			{
				actual:   str.New("...hello").FirstWord().String(),
				expected: "...hello",
			},
		},
	},
	{
		name: "Last",
		fixtures: []fixture{
			{
				actual:   str.New("golang").Last().String(),
				expected: "g",
			},
			{
				actual:   str.New("hello world").Last().String(),
				expected: "d",
			},
			{
				actual:   str.New("🔥 Go 🤩✨").Last().String(),
				expected: "✨",
			},
			{
				actual:   str.New("æøå").Last().String(),
				expected: "å",
			},
			{
				actual:   str.New("").Last().String(),
				expected: "",
			},
		},
	},
	{
		name: "LastWord",
		fixtures: []fixture{
			{
				actual:   str.New("golang").LastWord().String(),
				expected: "golang",
			},
			{
				actual:   str.New("hello world").LastWord().String(),
				expected: "world",
			},
			{
				actual:   str.New("    hello it    should   trim ").LastWord().String(),
				expected: "trim",
			},
			{
				actual:   str.New("æ ø å").LastWord().String(),
				expected: "å",
			},
			{
				actual:   str.New("").LastWord().String(),
				expected: "",
			},
			{
				actual:   str.New("...hello").LastWord().String(),
				expected: "...hello",
			},
		},
	},
	{
		name: "Length",
		fixtures: []fixture{
			{
				actual:   str.New("go🔥").Length(),
				expected: 3,
			},
			{
				actual:   str.New("golang").Length(),
				expected: 6,
			},
			{
				actual:   str.New("lorem ipsum dolor").Length(),
				expected: 17,
			},
			{
				actual:   str.New("你好世界").Length(),
				expected: 4,
			},
			{
				actual:   str.New("你好界").Length(),
				expected: 3,
			},
		},
	},
	{
		name: "Lower",
		fixtures: []fixture{
			{
				actual:   str.New("").Lower().String(),
				expected: "",
			},
			{
				actual:   str.New("✨").Lower().String(),
				expected: "✨",
			},
			{
				actual:   str.New("&").Lower().String(),
				expected: "&",
			},
			{
				actual:   str.New("Æ").Lower().String(),
				expected: "æ",
			},
			{
				actual:   str.New("TEST").Lower().String(),
				expected: "test",
			},
		},
	},
	{
		name: "IsEmpty",
		fixtures: []fixture{
			{
				actual:   str.New("golang").IsEmpty(),
				expected: false,
			},
			{
				actual:   str.New("").IsEmpty(),
				expected: true,
			},
			{
				actual:   str.New("   ").IsEmpty(),
				expected: false,
			},
		},
	},
	{
		name: "IsJson",
		fixtures: []fixture{
			{
				actual:   str.New(`{"key": "value"}`).IsJson(),
				expected: true,
			},
			{
				actual:   str.New(`{"key": "value", "nested": {"inner": 42}}`).IsJson(),
				expected: true,
			},
			{
				actual:   str.New("not a valid JSON").IsJson(),
				expected: false,
			},
			{
				actual:   str.New("").IsJson(),
				expected: false,
			},
		},
	},
	{
		name: "IsNotEmpty",
		fixtures: []fixture{
			{
				actual:   str.New("golang").IsNotEmpty(),
				expected: true,
			},
			{
				actual:   str.New("").IsNotEmpty(),
				expected: false,
			},
			{
				actual:   str.New("   ").IsNotEmpty(),
				expected: true,
			},
		},
	},
	{
		name: "IsUrl",
		fixtures: []fixture{
			{
				actual:   str.New("http://example.com").IsUrl(),
				expected: true,
			},
			{
				actual:   str.New("https://subdomain.example.com").IsUrl(),
				expected: true,
			},
			{
				actual:   str.New("https://example.com/path").IsUrl(),
				expected: true,
			},
			{
				actual:   str.New("https://example.com/path?query=value").IsUrl(),
				expected: true,
			},
			{
				actual:   str.New("https://example.com/path#section").IsUrl(),
				expected: true,
			},
			{
				actual:   str.New("https://user:password@example.com").IsUrl(),
				expected: true,
			},
			{
				actual:   str.New("https://example.com:8080").IsUrl(),
				expected: true,
			},
			{
				actual:   str.New("https://my-domain-example.com").IsUrl(),
				expected: true,
			},
			{
				actual:   str.New("https://my_domain_example.com").IsUrl(),
				expected: false,
			},
			{
				actual:   str.New("https://192.168.0.1").IsUrl(),
				expected: true,
			},
			{
				actual:   str.New("ftp://example.com").IsUrl(),
				expected: true,
			},
			{
				actual:   str.New("example.com").IsUrl(),
				expected: false,
			},
			{
				actual:   str.New("https://").IsUrl(),
				expected: false,
			},
			{
				actual:   str.New("https://example.com/").IsUrl(),
				expected: true,
			},
			{
				actual:   str.New("https://example!.com").IsUrl(),
				expected: false,
			},
			{
				actual:   str.New("https:/example.com").IsUrl(),
				expected: false,
			},
			{
				actual:   str.New("https//example.com").IsUrl(),
				expected: false,
			},
			{
				actual:   str.New("https://example .com").IsUrl(),
				expected: false,
			},
			{
				actual:   str.New("https://localhost:8080").IsUrl(),
				expected: true,
			},
			{
				actual:   str.New("https://example&com").IsUrl(),
				expected: false,
			},
			{
				actual:   str.New("192.168.0.0").IsUrl(),
				expected: false,
			},
			{
				actual:   str.New("http://192.168.0.0").IsUrl(),
				expected: true,
			},
			{
				actual:   str.New("https://192.168.0").IsUrl(),
				expected: true,
			},
			{
				actual:   str.New("https://🔥.rs").IsUrl(),
				expected: true,
			},
		},
	},
	{
		name: "Kebab",
		fixtures: []fixture{
			{
				actual:   str.New("").Kebab().String(),
				expected: "",
			},
			{
				actual:   str.New("go go go").Kebab().String(),
				expected: "go-go-go",
			},
			{
				actual:   str.New("GoLang").Kebab().String(),
				expected: "go-lang",
			},
			{
				actual:   str.New("goLang").Kebab().String(),
				expected: "go-lang",
			},
			{
				actual:   str.New("✨🔥✨🔥✨").Kebab().String(),
				expected: "✨🔥✨🔥✨",
			},
			{
				actual:   str.New("✨ 🔥 ✨ 🔥 ✨").Kebab().String(),
				expected: "✨-🔥-✨-🔥-✨",
			},
		},
	},
	{
		name: "Pascal",
		fixtures: []fixture{
			{
				actual:   str.New("").Pascal().String(),
				expected: "",
			},
			{
				actual:   str.New("go go go").Pascal().String(),
				expected: "GoGoGo",
			},
			{
				actual:   str.New("hello world").Pascal().String(),
				expected: "HelloWorld",
			},
			{
				actual:   str.New("Go-lang").Pascal().String(),
				expected: "Go-lang",
			},
			{
				actual:   str.New("goLang").Pascal().String(),
				expected: "GoLang",
			},
			{
				actual:   str.New("✨🔥✨🔥✨").Pascal().String(),
				expected: "✨🔥✨🔥✨",
			},
			{
				actual:   str.New("✨ 🔥 ✨ 🔥 ✨").Pascal().String(),
				expected: "✨🔥✨🔥✨",
			},
		},
	},
	{
		name: "Prepend",
		fixtures: []fixture{
			{
				actual:   str.New("♥️ is").Prepend("golang").String(),
				expected: "golang♥️ is",
			},
			{
				actual:   str.New("!@#$%^&*()").Prepend("hello").String(),
				expected: "hello!@#$%^&*()",
			},
			{
				actual:   str.New("").Prepend("øæå").String(),
				expected: "øæå",
			},
			{
				actual:   str.New("🚀").Prepend("test").String(),
				expected: "test🚀",
			},
		},
	},
	{
		name: "Repeat",
		fixtures: []fixture{
			{
				actual:   str.New("golang").Repeat(3).String(),
				expected: "golanggolanggolang",
			},
			{
				actual:   str.New("golang").Repeat(0).String(),
				expected: "",
			},
			{
				actual:   str.New("lorem ipsum").Repeat(2).String(),
				expected: "lorem ipsumlorem ipsum",
			},
			{
				actual:   str.New("").Repeat(5).String(),
				expected: "",
			},
		},
	},
	{
		name: "Remove",
		fixtures: []fixture{
			{
				actual:   str.New("i like c++").Remove("c++").String(),
				expected: "i like ",
			},
			{
				actual:   str.New("rocket 🚀 man").Remove("🚀").String(),
				expected: "rocket  man",
			},
			{
				actual:   str.New("go, go, go!").Remove("go").String(),
				expected: ", , !",
			},
			{
				actual:   str.New("你好世界").Remove("你好").String(),
				expected: "世界",
			},
			{
				actual:   str.New("").Remove("go").String(),
				expected: "",
			},
		},
	},
	{
		name: "Replace",
		fixtures: []fixture{
			{
				actual:   str.New("i like c++").Replace("c++", "go").String(),
				expected: "i like go",
			},
			{
				actual:   str.New("rocket 🚀 man").Replace("🚀", "🌎").String(),
				expected: "rocket 🌎 man",
			},
			{
				actual:   str.New("go, go, go!").Replace("go", "Go").String(),
				expected: "Go, Go, Go!",
			},
			{
				actual:   str.New("你好世界").Replace("你好", "再见").String(),
				expected: "再见世界",
			},
			{
				actual:   str.New("你好世界").Replace("你好", "").String(),
				expected: "世界",
			},
		},
	},
	{
		name: "Screaming",
		fixtures: []fixture{
			{
				actual:   str.New("go go go").Screaming().String(),
				expected: "GO_GO_GO",
			},
			{
				actual:   str.New("GoLang").Screaming().String(),
				expected: "GO_LANG",
			},
			{
				actual:   str.New("ALREADY_SCREAMING").Screaming().String(),
				expected: "ALREADY_SCREAMING",
			},
			{
				actual:   str.New("").Screaming().String(),
				expected: "",
			},
			{
				actual:   str.New("goLang").Screaming().String(),
				expected: "GO_LANG",
			},
			{
				actual:   str.New("✨🔥✨🔥✨").Screaming().String(),
				expected: "✨🔥✨🔥✨",
			},
			{
				actual:   str.New("✨ 🔥 ✨ 🔥 ✨").Screaming().String(),
				expected: "✨_🔥_✨_🔥_✨",
			},
		},
	},
	{
		name: "Slug",
		fixtures: []fixture{
			{
				actual:   str.New("  ❤️ go ❤️  ").Slug("_").String(),
				expected: "go",
			},
			{
				actual:   str.New("  go ❤️  ").Slug("_").String(),
				expected: "go",
			},
			{
				actual:   str.New("GoLang").Slug("_").String(),
				expected: "golang",
			},
			{
				actual:   str.New("foo_@_bar-baz!").Slug("_").String(),
				expected: "foo_bar_baz",
			},
			{
				actual:   str.New("foo_@_bar-baz!").Slug("-").String(),
				expected: "foo-bar-baz",
			},
			{
				actual:   str.New("foo_@_bar-baz!").Slug("/").String(),
				expected: "foo/bar/baz",
			},
			{
				actual:   str.New("").Slug("_").String(),
				expected: "",
			},
			{
				actual:   str.New("goLang").Slug("_").String(),
				expected: "golang",
			},
			{
				actual:   str.New("✨🔥✨🔥✨").Slug("_").String(),
				expected: "",
			},
			{
				actual:   str.New("✨ 🔥 ✨ 🔥 ✨").Slug("_").String(),
				expected: "",
			},
			{
				actual:   str.New("__hello__world__").Slug("_").String(),
				expected: "hello_world",
			},
			{
				actual:   str.New("__hello__world__").Slug("-").String(),
				expected: "hello-world",
			},
		},
	},
	{
		name: "Snake",
		fixtures: []fixture{
			{
				actual:   str.New("go go go").Snake().String(),
				expected: "go_go_go",
			},
			{
				actual:   str.New("GoLang").Snake().String(),
				expected: "go_lang",
			},
			{
				actual:   str.New("already_snake_cased").Snake().String(),
				expected: "already_snake_cased",
			},
			{
				actual:   str.New("").Snake().String(),
				expected: "",
			},
			{
				actual:   str.New("goLang").Snake().String(),
				expected: "go_lang",
			},
			{
				actual:   str.New("✨🔥✨🔥✨").Snake().String(),
				expected: "✨🔥✨🔥✨",
			},
			{
				actual:   str.New("✨ 🔥 ✨ 🔥 ✨").Snake().String(),
				expected: "✨_🔥_✨_🔥_✨",
			},
		},
	},
	{
		name: "Split",
		fixtures: []fixture{
			{
				actual:   str.New("go").Split(""),
				expected: []string{"g", "o"},
			},
			{
				actual:   str.New("go go go").Split(" "),
				expected: []string{"go", "go", "go"},
			},
			{
				actual:   str.New("golang").Split(" "),
				expected: []string{"golang"},
			},
			{
				actual:   str.New("lorem ipsum").Split("_"),
				expected: []string{"lorem ipsum"},
			},
			{
				actual:   str.New("").Split(""),
				expected: []string{},
			},
			{
				actual:   str.New("    go  go  go").Split(""),
				expected: []string{" ", " ", " ", " ", "g", "o", " ", " ", "g", "o", " ", " ", "g", "o"},
			},
			{
				actual:   str.New("✨🔥✨🔥✨").Split("🔥"),
				expected: []string{"✨", "✨", "✨"},
			},
		},
	},
	{
		name: "StartsWith",
		fixtures: []fixture{
			{
				actual:   str.New("golang").StartsWith(""),
				expected: false,
			},
			{
				actual:   str.New("golang").StartsWith("go"),
				expected: true,
			},
			{
				actual:   str.New(" go ").StartsWith(" "),
				expected: true,
			},
			{
				actual:   str.New("🔥🔥").StartsWith("🔥"),
				expected: true,
			},
			{
				actual:   str.New("✨🔥✨🔥✨").StartsWith("🔥"),
				expected: false,
			},
		},
	},
	{
		name: "String",
		fixtures: []fixture{
			{
				actual:   str.New("go").Capitalize().String(),
				expected: "Go",
			},
			{
				actual:   str.New("").Lower().String(),
				expected: "",
			},
		},
	},
	{
		name: "Swap",
		fixtures: []fixture{
			{
				actual: str.New("I love Rust").
					Swap(
						map[string]string{
							"Rust": "Go",
						}).
					String(),
				expected: "I love Go",
			},
			{
				actual: str.New("Rust. Rust! RUST!!").
					Swap(
						map[string]string{
							"Rust": "Go",
						}).
					String(),
				expected: "Go. Go! RUST!!",
			},
			{
				actual: str.New("Rust. Rust! RUST!!").
					Swap(
						map[string]string{
							"Rust": "Go",
							"RUST": "Go",
						}).
					String(),
				expected: "Go. Go! Go!!",
			},
			{
				actual: str.New("I ❤️ U").
					Swap(
						map[string]string{
							"❤️": "🤬",
							"I":  "You",
							"U":  "Me",
						}).
					String(),
				expected: "You 🤬 Me",
			},
			{
				actual: str.New("BLÅBÆRSYLTETØY blåbærsyltetøy").
					Swap(
						map[string]string{
							"Æ": "",
							"Ø": "",
							"Å": "",
							"æ": "",
							"ø": "",
							"å": "",
						}).
					String(),
				expected: "BLBRSYLTETY blbrsyltety",
			},
			{
				actual: str.New("").
					Swap(
						map[string]string{
							"Æ": "",
							"Ø": "",
							"Å": "",
							"æ": "",
							"ø": "",
							"å": "",
						}).
					String(),
				expected: "",
			},
		},
	},
	{
		name: "Title",
		fixtures: []fixture{
			{
				actual:   str.New("go øl åre ære").Title().String(),
				expected: "Go Øl Åre Ære",
			},
			{
				actual:   str.New("go ✨ go").Title().String(),
				expected: "Go ✨ Go",
			},
			{
				actual:   str.New("_ _ _").Title().String(),
				expected: "_ _ _",
			},
			{
				actual:   str.New("").Title().String(),
				expected: "",
			},
		},
	},
	{
		name: "TrimAll",
		fixtures: []fixture{
			{
				actual:   str.New(" Hello,   World  ! ").TrimAll().String(),
				expected: "Hello, World !",
			},
			{
				actual:   str.New("Hello,\tWorld ! ").TrimAll().String(),
				expected: "Hello, World !",
			},
			{
				actual:   str.New(" \t\n\t Hello,\tWorld\n!\n\t").TrimAll().String(),
				expected: "Hello, World !",
			},
		},
	},
	{
		name: "TrimBoth",
		fixtures: []fixture{
			{
				actual:   str.New(" Hello,   World  ! ").TrimBoth().String(),
				expected: "Hello,   World  !",
			},
			{
				actual:   str.New("Hello,\tWorld ! ").TrimBoth().String(),
				expected: "Hello,\tWorld !",
			},
			{
				actual:   str.New(" \t\n\t Hello,\tWorld\n!\n\t").TrimBoth().String(),
				expected: "Hello,\tWorld\n!",
			},
		},
	},
	{
		name: "TrimLeft",
		fixtures: []fixture{
			{
				actual:   str.New(" Hello,   World  ! ").TrimLeft().String(),
				expected: "Hello,   World  ! ",
			},
			{
				actual:   str.New("Hello,\tWorld ! ").TrimLeft().String(),
				expected: "Hello,\tWorld ! ",
			},
			{
				actual:   str.New(" \t\n\t Hello,\tWorld\n!\n\t").TrimLeft().String(),
				expected: "Hello,\tWorld\n!\n\t",
			},
		},
	},
	{
		name: "TrimRight",
		fixtures: []fixture{
			{
				actual:   str.New(" Hello,   World  ! ").TrimRight().String(),
				expected: " Hello,   World  !",
			},
			{
				actual:   str.New("Hello,\tWorld ! ").TrimRight().String(),
				expected: "Hello,\tWorld !",
			},
			{
				actual:   str.New(" \t\n\t Hello,\tWorld\n!\n\t").TrimRight().String(),
				expected: " \t\n\t Hello,\tWorld\n!",
			},
		},
	},
	{
		name: "Upper",
		fixtures: []fixture{
			{
				actual:   str.New("").Upper().String(),
				expected: "",
			},
			{
				actual:   str.New("✨").Upper().String(),
				expected: "✨",
			},
			{
				actual:   str.New("&").Upper().String(),
				expected: "&",
			},
			{
				actual:   str.New("æ").Upper().String(),
				expected: "Æ",
			},
			{
				actual:   str.New("test").Upper().String(),
				expected: "TEST",
			},
		},
	},
	{
		name: "When",
		fixtures: []fixture{
			{
				actual: str.New("go").
					When(true, func(s *str.Str) *str.Str {
						return s.Append("lang")
					}).
					String(),
				expected: "golang",
			},
			{
				actual: str.New("go").
					When(false, func(s *str.Str) *str.Str {
						return s.Append("lang")
					}).
					String(),
				expected: "go",
			},
			{
				actual: str.New("").
					When(false, func(s *str.Str) *str.Str {
						return s.Append("lang")
					}).
					String(),
				expected: "",
			},
		},
	},
	{
		name: "WhenContains",
		fixtures: []fixture{
			{
				actual: str.New("i like go").
					WhenContains("go", func(s *str.Str) *str.Str {
						return s.Append("lang")
					}).
					String(),
				expected: "i like golang",
			},
			{
				actual: str.New("i like go").
					WhenContains("gox", func(s *str.Str) *str.Str {
						return s.Append("lang")
					}).
					String(),
				expected: "i like go",
			},
			{
				actual: str.New("").
					WhenContains("go", func(s *str.Str) *str.Str {
						return s.Append("lang")
					}).
					String(),
				expected: "",
			},
		},
	},
	{
		name: "WhenContainsAny",
		fixtures: []fixture{
			{
				actual: str.New("i like go").
					WhenContainsAny([]string{"go", "c++"}, func(s *str.Str) *str.Str {
						return s.Append("lang")
					}).
					String(),
				expected: "i like golang",
			},
			{
				actual: str.New("i like go").
					WhenContainsAny([]string{"a", "b", "c"}, func(s *str.Str) *str.Str {
						return s.Append("lang")
					}).
					String(),
				expected: "i like go",
			},
			{
				actual: str.New("").
					WhenContainsAny([]string{"go", "lang"}, func(s *str.Str) *str.Str {
						return s.Append("lang")
					}).
					String(),
				expected: "",
			},
		},
	},
	{
		name: "WhenContainsAll",
		fixtures: []fixture{
			{
				actual: str.New("i like go not c++").
					WhenContainsAll([]string{"go", "c++"}, func(s *str.Str) *str.Str {
						return s.Append("!")
					}).
					String(),
				expected: "i like go not c++!",
			},
			{
				actual: str.New("i like go").
					WhenContainsAll([]string{"a", "b", "c"}, func(s *str.Str) *str.Str {
						return s.Append("lang")
					}).
					String(),
				expected: "i like go",
			},
			{
				actual: str.New("").
					WhenContainsAll([]string{"go", "lang"}, func(s *str.Str) *str.Str {
						return s.Append("lang")
					}).
					String(),
				expected: "",
			},
		},
	},
	{
		name: "WhenEndsWith",
		fixtures: []fixture{
			{
				actual: str.New("i like go").
					WhenEndsWith("go", func(s *str.Str) *str.Str {
						return s.Append("lang")
					}).
					String(),
				expected: "i like golang",
			},
			{
				actual: str.New("i like go").
					WhenEndsWith("gox", func(s *str.Str) *str.Str {
						return s.Append("lang")
					}).
					String(),
				expected: "i like go",
			},
			{
				actual: str.New("").
					WhenEndsWith("go", func(s *str.Str) *str.Str {
						return s.Append("lang")
					}).
					String(),
				expected: "",
			},
		},
	},
	{
		name: "WhenEmpty",
		fixtures: []fixture{
			{
				actual: str.New("go").
					WhenEmpty(func(s *str.Str) *str.Str {
						return s.Append("lang")
					}).
					String(),
				expected: "go",
			},
			{
				actual: str.New("").
					WhenEmpty(func(s *str.Str) *str.Str {
						return s.Append("golang")
					}).
					String(),
				expected: "golang",
			},
		},
	},
	{
		name: "WhenExactly",
		fixtures: []fixture{
			{
				actual: str.New("go").
					WhenExactly("go", func(s *str.Str) *str.Str {
						return s.Append("lang")
					}).
					String(),
				expected: "golang",
			},
			{
				actual: str.New("").
					WhenExactly("go", func(s *str.Str) *str.Str {
						return s.Append("golang")
					}).
					String(),
				expected: "",
			},
			{
				actual: str.New("🚀").
					WhenExactly("🚀", func(s *str.Str) *str.Str {
						return s.Append("✨")
					}).
					String(),
				expected: "🚀✨",
			},
		},
	},
	{
		name: "WhenFunc",
		fixtures: []fixture{
			{
				actual: str.New("hello world").
					WhenFunc(func(s *str.Str) bool {
						return s.FirstWord().Exactly("hello")
					}, func(s *str.Str) *str.Str {
						return s.Append("!")
					}).
					String(),
				expected: "hello world!",
			},
		},
	},
	{
		name: "WhenNotEmpty",
		fixtures: []fixture{
			{
				actual: str.New("go").
					WhenNotEmpty(func(s *str.Str) *str.Str {
						return s.Append("lang")
					}).
					String(),
				expected: "golang",
			},
			{
				actual: str.New("").
					WhenNotEmpty(func(s *str.Str) *str.Str {
						return s.Append("golang")
					}).
					String(),
				expected: "",
			},
		},
	},
	{
		name: "WhenNotExactly",
		fixtures: []fixture{
			{
				actual: str.New("go").
					WhenNotExactly("go", func(s *str.Str) *str.Str {
						return s.Append("lang")
					}).
					String(),
				expected: "go",
			},
			{
				actual: str.New("").
					WhenNotExactly("go", func(s *str.Str) *str.Str {
						return s.Append("golang")
					}).
					String(),
				expected: "golang",
			},
			{
				actual: str.New("🚀").
					WhenNotExactly("🚀", func(s *str.Str) *str.Str {
						return s.Append("✨")
					}).
					String(),
				expected: "🚀",
			},
		},
	},
	{
		name: "WhenStartsWith",
		fixtures: []fixture{
			{
				actual: str.New("go").
					WhenStartsWith("go", func(s *str.Str) *str.Str {
						return s.Append("lang")
					}).
					String(),
				expected: "golang",
			},
			{
				actual: str.New("go").
					WhenStartsWith("gox", func(s *str.Str) *str.Str {
						return s.Append("lang")
					}).
					String(),
				expected: "go",
			},
			{
				actual: str.New("").
					WhenStartsWith("go", func(s *str.Str) *str.Str {
						return s.Append("lang")
					}).
					String(),
				expected: "",
			},
		},
	},
	{
		name: "Words",
		fixtures: []fixture{
			{
				actual:   str.New("    go  go  go").Words(),
				expected: []string{"go", "go", "go"},
			},
			{
				actual:   str.New("go").Words(),
				expected: []string{"go"},
			},
			{
				actual:   str.New("-_%$!!$&(=!)").Words(),
				expected: []string{"-_%$!!$&(=!)"},
			},
			{
				actual:   str.New("go go go").Words(),
				expected: []string{"go", "go", "go"},
			},
			{
				actual:   str.New("golang").Words(),
				expected: []string{"golang"},
			},
			{
				actual:   str.New("lorem ipsum").Words(),
				expected: []string{"lorem", "ipsum"},
			},
			{
				actual:   str.New("").Words(),
				expected: []string{},
			},
			{
				actual:   str.New("✨🔥✨🔥✨").Words(),
				expected: []string{"✨🔥✨🔥✨"},
			},
		},
	},
}

func TestStr(t *testing.T) {
	for _, tc := range testCases {
		for _, f := range tc.fixtures {
			t.Run(tc.name, func(t *testing.T) {
				if !reflect.DeepEqual(f.actual, f.expected) {
					t.Fatalf("Expected: %q, Actual: %q", f.expected, f.actual)
				}
			})
		}
	}
}

func TestDocs(t *testing.T) {
	cmd := exec.Command("sh", "-c", "cd docs/; go run docs.go")
	_, err := cmd.Output()
	if err != nil {
		// if there was any error, print it here
		fmt.Println("could not run command: ", err)
	}

	log.Println("Docs generated")
}

func TestEveryPublicMethod(t *testing.T) {
	var testedFunctions = make(map[string]bool)
	for _, tc := range testCases {
		testedFunctions[tc.name] = true
	}

	var implementedFunctions = make(map[string]bool)
	content, err := os.ReadFile("str.go")

	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")

	for _, line := range lines {
		if strings.Contains(line, "func (s *Str)") {
			fn := str.New(line).Between("func (s *Str)", "(").TrimBoth()
			firstChar := fn.FirstWord().String()

			if firstChar == strings.ToUpper(firstChar) {
				implementedFunctions[fn.String()] = true
			}
		}
	}

	var missingFunctions []string
	for key := range implementedFunctions {
		if !testedFunctions[key] {
			missingFunctions = append(missingFunctions, key)
		}
	}

	if len(missingFunctions) > 0 {
		t.Fatalf("Missing functions: %v", missingFunctions)
	}
}
