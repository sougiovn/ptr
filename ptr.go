package ptr

// generic ptr
func Ptr[T any](v T) *T {
	return &v
}

func PtrSlice[T any](v []T) []*T {
	p := make([]*T, len(v))
	for i := range v {
		p[i] = &v[i]
	}
	return p
}

func PtrMap[K comparable, T any](v map[K]T) map[K]*T {
	p := make(map[K]*T, len(v))
	for k, v := range v {
		p[k] = &v
	}
	return p
}

// generic value
func Value[T any](p *T) T {
	if p != nil {
		return *p
	}
	var v T
	return v
}

func ValueSlice[T any](p []*T) []T {
	v := make([]T, len(p))
	for i := range p {
		val := p[i]
		if val != nil {
			v[i] = *val
		}
	}
	return v
}

func ValueMap[K comparable, T any](p map[K]*T) map[K]T {
	v := make(map[K]T, len(p))
	for key, val := range p {
		if val != nil {
			v[key] = *val
		}
	}
	return v
}
