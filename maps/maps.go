package maps

import (
	"sort"
)

// Keys returns the keys of the map m.
// The keys will be in an indeterminate order.
func Keys[M ~map[K]V, K comparable, V any](m M) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

func Keys_ordered[M ~map[K]V, K int64, V any](m M) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return r
}

// Values returns the values of the map m.
// The values will be in an indeterminate order.
func Values[M ~map[K]V, K comparable, V any](m M) []V {
	r := make([]V, 0, len(m))
	for _, v := range m {
		r = append(r, v)
	}
	return r
}

func Values2[M ~map[int64]V, V any](m M) []V {
	r := make([]V, len(m))
	for i := int64(0); i < int64(len(m)); i++ {
		r[i] = m[i]
	}
	return r
}

func Pop[T any](stack *[]T) {
	if len(*stack) == 0 {
		return
	}
	*stack = (*stack)[:len(*stack)-1]
}
func Pop_front[T any](stack *[]T) T {
	if len(*stack) == 0 {
		var NULL T
		return NULL
	} else {
		poped := (*stack)[0]
		*stack = (*stack)[1:]
		return poped
	}

}
