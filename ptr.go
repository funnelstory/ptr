package ptr

// Ptr returns a pointer to the provided value.
func Ptr[T any](v T) *T {
	return &v
}

// Deref returns the value pointed to by p, or the zero value of T if p is nil.
func Deref[T any](p *T) T {
	if p == nil {
		var zero T
		return zero
	}
	return *p
}
