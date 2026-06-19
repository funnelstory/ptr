# 🧠 ptr — Pointer to Anything

Generic helpers for pointerizing literals and safely dereferencing in Go.

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

That’s it. Two small functions. Fully generic. Reusable anywhere. Test-covered. Type-safe.

---

## 📦 Installation

```sh
go get github.com/funnelstory/ptr
```

Or if you're using modules, just import it and `go mod tidy` will do the rest.

---

## 🧪 Usage

The `ptr` package exports:

```go
func Ptr[T any](v T) *T
func Deref[T any](p *T) T
```

- **Ptr** — turn any value into a pointer.
- **Deref** — get the value from a pointer, or the zero value of `T` if the pointer is nil (no panic).

### Example:

```go
package main

import (
	"fmt"
	"github.com/funnelstory/ptr"
)

func main() {
	s := ptr.Ptr("golang")
	fmt.Println(ptr.Deref(s)) // Output: golang

	// Safe with nil:
	var p *string = nil
	fmt.Println(ptr.Deref(p)) // Output: (empty string, no panic)
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

### Deref examples:

```go
ptr.Deref(ptr.Ptr(42))       // 42
ptr.Deref((*int)(nil))       // 0  (zero value, no panic)
ptr.Deref(ptr.Ptr("hi"))     // "hi"
```

---

## 💼 Why?

- Avoid boilerplate `tmp := val; &tmp`
- Helpful in:
  - JSON/YAML marshaling with optional `*T` fields
  - Config structs
  - Unit testing
  - Working with APIs that use pointer types for optionality
- **Deref**: safely read optional `*T` fields without nil checks or panics
- Tiny and dependency-free
- Idiomatic: one package, one purpose

---

## 🔍 Project Layout

```text
ptr/                       # repo root = module root = package dir
├── ptr.go                 # Ptr + Deref
├── ptr_test.go            # Tests covering many types
├── go.mod
├── Makefile               # test, lint, etc.
├── .github/
│   └── workflows/
│       └── ci.yml         # GitHub Actions for lint + test
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
- Benchmarks

Open an issue or PR — and please include tests!
