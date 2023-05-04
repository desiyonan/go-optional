package optional

type Optional[T any] interface {
	Get() (T, bool)
	Or(T) Optional[T]
	GetOrZero() T
	GetOrElse(T) T
	IsZero() bool
}

// Map convert an Option of type T to an Option of type K.
//func Map[T, K any](mapper func(t T) K, opt Option[T]) Option[K] {
//	if opt.IsZero() {
//		return None[K]()
//	}
//	value := opt.Get()
//	return Some(mapper(value))
//}

func Of[T any](t T) Optional[T] {
	return Some(t)
}

func Any[T any](t T, others ...T) Optional[T] {
	opt := Some(t)
	if !opt.IsZero() {
		return opt
	}

	for _, o := range others {
		opt = opt.Or(o)
		if !opt.IsZero() {
			return opt
		}
	}

	return opt
}

func IsZero[T any](t T) bool {
	return Of(t).IsZero()
}
