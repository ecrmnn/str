---
title: "str - A fluent string manipulation library in Go"
---

# ![str](./assets/str.png)

## A fluent string manipulation library in Go

### Installation

```bash
go get -u github.com/ecrmnn/str
```

### Examples

```go
s := "String Manipulation"

str.New(s).Snake().Prepend("dir/").Append(".md").String()

// dir/string_manipulation.md
```

```go
s := "What I think of C++? I try to avoid C++..."

m := map[string]string{
  "C++":    "Go",
  "avoid":  "love",
  "try to": "",
  "...":    "!",
}

str.New(s).Swap(m).Append(" 🤩").TrimAll().String()

// What I think of Go? I love Go! 🤩
```
