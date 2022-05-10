package goint

// To returns a pointer to a variable
func To[T any](t T) *T {
	return &t
}

// To returns a variable from a pointer or a default value if pointer is nil
func From[T any](t *T, def T) T {
	if t != nil {
		return *t
	}
	return def
}
