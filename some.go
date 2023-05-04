package optional

import "reflect"

type some[T any] struct {
	val T
}

func (s some[T]) IsZero() bool {
	var t T
	return reflect.DeepEqual(s.val, t)
}

func (s some[T]) GetValue() T {
	return s.val
}

func (s some[T]) Get() (T, bool) {
	return s.val, s.IsZero()
}

// Or 只有当前为零值时，则取传入的值，创建新的Optional返回，原Optional不改变
func (s some[T]) Or(t T) Optional[T] {
	if s.IsZero() {
		return Some(t)
	}
	return s
}

// OrElse 只有当前为零值时，则返回传入的值
func (s some[T]) OrElse(t T) T {
	if s.IsZero() {
		return t
	}
	return s.val
}

// OrGet 只有当前为零值时，则执行传入的GetFun
func (s some[T]) OrGet(getter GetFun[T]) T {
	if s.IsZero() {
		return getter()
	}
	return s.val
}

// Map 如果当前值不为零值，则映射到另外的Optional
//func (s some[T]) Map(mapper MapFun[T, any]) Optional[any] {
//	if s.IsZero() {
//		var z any
//		return Some(z)
//	}
//	return Some(mapper(s.val))
//}

func Some[T any](val T) Optional[T] {
	return some[T]{
		val: val,
	}
}
