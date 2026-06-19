package ptr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPtr(t *testing.T) {
	t.Parallel()

	t.Run("int", func(t *testing.T) {
		t.Parallel()
		v := 42
		p := Ptr(v)
		assert.NotNil(t, p)
		assert.Equal(t, v, *p)
	})

	t.Run("string", func(t *testing.T) {
		t.Parallel()
		s := "hello"
		p := Ptr(s)
		assert.NotNil(t, p)
		assert.Equal(t, s, *p)
	})

	t.Run("empty string", func(t *testing.T) {
		t.Parallel()
		s := ""
		p := Ptr(s)
		assert.NotNil(t, p)
		assert.Equal(t, s, *p)
	})

	t.Run("zero int", func(t *testing.T) {
		t.Parallel()
		var i int
		p := Ptr(i)
		assert.NotNil(t, p)
		assert.Equal(t, 0, *p)
	})

	t.Run("bool", func(t *testing.T) {
		t.Parallel()
		b := true
		p := Ptr(b)
		assert.NotNil(t, p)
		assert.Equal(t, b, *p)
	})

	t.Run("float64", func(t *testing.T) {
		t.Parallel()
		f := 3.1415
		p := Ptr(f)
		assert.NotNil(t, p)
		assert.Equal(t, f, *p)
	})

	t.Run("struct", func(t *testing.T) {
		t.Parallel()
		type foo struct {
			A int
			B string
		}
		val := foo{A: 10, B: "bar"}
		p := Ptr(val)
		assert.NotNil(t, p)
		assert.Equal(t, val, *p)
	})

	t.Run("slice", func(t *testing.T) {
		t.Parallel()
		s := []string{"a", "b"}
		p := Ptr(s)
		assert.NotNil(t, p)
		assert.Equal(t, s, *p)
	})

	t.Run("map", func(t *testing.T) {
		t.Parallel()
		m := map[string]int{"x": 1}
		p := Ptr(m)
		assert.NotNil(t, p)
		assert.Equal(t, m, *p)
	})

	t.Run("nilable pointer", func(t *testing.T) {
		t.Parallel()
		var x *int = nil
		p := Ptr(x)
		assert.NotNil(t, p)
		assert.Equal(t, x, *p)
	})

	t.Run("nested pointer", func(t *testing.T) {
		t.Parallel()
		x := 5
		px := &x
		p := Ptr(px)
		assert.NotNil(t, p)
		assert.Equal(t, px, *p)
		assert.Equal(t, 5, **p)
	})

	t.Run("func", func(t *testing.T) {
		t.Parallel()
		fn := func() string { return "ok" }
		p := Ptr(fn)
		assert.NotNil(t, p)
		assert.Equal(t, "ok", (*p)())
	})

	t.Run("interface", func(t *testing.T) {
		t.Parallel()
		var i interface{} = 123
		p := Ptr(i)
		assert.NotNil(t, p)
		assert.Equal(t, 123, *p)
	})
}

func TestDeref(t *testing.T) {
	t.Parallel()

	t.Run("nil returns zero value", func(t *testing.T) {
		t.Parallel()
		var p *int = nil
		assert.Equal(t, 0, Deref(p))
	})

	t.Run("nil string returns empty string", func(t *testing.T) {
		t.Parallel()
		var p *string = nil
		assert.Equal(t, "", Deref(p))
	})

	t.Run("nil bool returns false", func(t *testing.T) {
		t.Parallel()
		var p *bool = nil
		assert.Equal(t, false, Deref(p))
	})

	t.Run("non-nil returns value", func(t *testing.T) {
		t.Parallel()
		v := 42
		p := &v
		assert.Equal(t, 42, Deref(p))
	})

	t.Run("string", func(t *testing.T) {
		t.Parallel()
		s := "hello"
		p := &s
		assert.Equal(t, "hello", Deref(p))
	})

	t.Run("struct", func(t *testing.T) {
		t.Parallel()
		type foo struct {
			A int
			B string
		}
		val := foo{A: 10, B: "bar"}
		p := &val
		assert.Equal(t, val, Deref(p))
	})

	t.Run("nil struct pointer returns zero struct", func(t *testing.T) {
		t.Parallel()
		type foo struct {
			A int
			B string
		}
		var p *foo = nil
		got := Deref(p)
		assert.Equal(t, 0, got.A)
		assert.Equal(t, "", got.B)
	})

	t.Run("roundtrip Ptr then Deref", func(t *testing.T) {
		t.Parallel()
		orig := 123
		assert.Equal(t, orig, Deref(Ptr(orig)))
	})
}
