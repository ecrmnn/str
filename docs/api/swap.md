# Swap `*Str`

`Swap` returns a new instance of Str where all occurrences of keys are replaces with their corresponding values.

### Signature

```go
func (s *Str) Swap(pairs map[string]string) *Str
```

### Examples

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
