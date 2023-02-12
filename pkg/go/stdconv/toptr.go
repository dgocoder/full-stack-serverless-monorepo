package stdconv

func Ptr[T any](t T) *T { return &t }

func PtrSlice[T any](t []T) (l []*T) {
	for i := range t {
		l = append(l, &t[i])
	}

	return
}
