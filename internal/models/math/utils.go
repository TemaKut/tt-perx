package mathmodels

func toPtr[T any](v T) *T {
	return &v
}
