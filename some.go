package optional

type some[T any] struct {
	val T
}

func (s some[T]) Get() (T, bool) {
	return s.val, s.IsZero()
}

func (s some[T]) Or(t T) Optional[T] {
	if s.IsZero() {
		return Some(t)
	}
	return s
}

func (s some[T]) GetOrElse(t T) T {
	if !s.IsZero() {
		return s.val
	}
	return t
}

func (s some[T]) GetOrZero() T {
	var t T
	return s.GetOrElse(t)
}

func (s some[T]) IsZero() bool {
	var t T
	//return s.val == s.val
	return s == some[T]{val: t}
}

func Some[T any](val T) Optional[T] {
	return some[T]{
		val: val,
	}
}
