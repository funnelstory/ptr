# 🧠 ptr — Pointer to Anything

The only generic function you need for pointerizing literals in Go.

---

## ✨ What is `ptr`?

Go is amazing, but it makes one thing surprisingly awkward:

```go
// You can't do this:
x := &"hello" // 🚫 compile error
```

You either:
- Make a temp variable like a caveman, or
- Write the same boilerplate helper in every repo.

Let’s fix that:

```go
import "github.com/funnelstory/ptr"

x := ptr.Ptr("hello") // ✅ *string
```

That’s it. One function. Fully generic. Reusable anywhere. Test-covered. Type-safe.

---

## 📦 Installation

```sh
go get github.com/funnelstory/ptr
```

Or if you're using modules, just import it and `go mod tidy` will do the rest.

---

## 🧪 Usage

The `ptr` package exports exactly one thing:

```go
func Ptr[T any](v T) *T
```

### Example:

```go
package main

import (
	"fmt"
	"github.com/funnelstory/ptr"
)

func main() {
	s := ptr.Ptr("golang")
	fmt.Println(*s) // Output: golang
}
```

### Works with everything:

```go
ptr.Ptr(42)                   // *int
ptr.Ptr(true)                 // *bool
ptr.Ptr(3.14)                 // *float64
ptr.Ptr([]string{"a", "b"})   // *[]string
ptr.Ptr(map[string]int{})     // *map[string]int
ptr.Ptr(func() {})            // *func()
ptr.Ptr(nil)                  // *any
```

---

## 💼 Why?

- Avoid boilerplate `tmp := val; &tmp`
- Helpful in:
  - JSON/YAML marshaling with optional `*T` fields
  - Config structs
  - Unit testing
  - Working with APIs that use pointer types for optionality
- Tiny and dependency-free
- Idiomatic: one package, one purpose

---

## 🔍 Project Layout

```text
ptr/
├── ptr/
│   └── ptr.go            # <- The one true function
│   └── ptr_test.go       # <- Tests covering many types
├── go.mod
├── Makefile              # <- test, lint, etc.
├── .github/
│   └── workflows/
│       └── ci.yml        # <- GitHub Actions for lint + test
```

---

## 🧪 Running Tests

```bash
go test ./...
```

Or use the Makefile:

```bash
make test
```

Tests are parallelized and include:

- Primitives
- Structs, maps, slices
- Nested pointers
- Interface and function values
- Nilables

---

## 🧠 Compatibility

- Requires Go **1.18+** (uses generics)
- Works with all built-in and user-defined types
- 100% test coverage

---

## 🚀 Contributing

This is a micro-library, but you’re welcome to contribute:

- Better docs/examples
- More weird test cases
- A `Deref[T]` companion (?)
- Benchmarks

Open an issue or PR — and please include tests!

```
