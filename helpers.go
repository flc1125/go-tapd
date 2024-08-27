package tapd

// Ptr returns a pointer to the value.
func Ptr[T any](v T) *T {
	return &v
}
