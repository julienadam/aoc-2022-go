package util

import "golang.org/x/exp/constraints"

type Range[T constraints.Ordered] struct {
	Start T
	End   T
}

func (m Range[T]) Overlaps(o Range[T]) bool {
	return (m.Start <= o.End) && (m.End >= o.Start)
}

func (m Range[T]) Contains(o Range[T]) bool {
	return o.Start >= m.Start && o.End <= m.End
}

func (m Range[T]) ContainsOrIsContainedIn(o Range[T]) bool {
	return m.Contains(o) || o.Contains(m)
}

func MakeRange[T constraints.Ordered](a T, b T) Range[T] {
	return Range[T]{a, b}
}
