package optional

// MapFun 映射S类型的值到R类型的值
// @type S 源类型
// @type R 结果类型
type MapFun[S any, R any] func(val S) R

type GetFun[T any] func() T

type Optional[T any] interface {
	// IsZero 是否为零值
	IsZero() bool
	// GetValue 获取值，配合IsZero，如果为零值依旧返回
	GetValue() T
	// Get 获取值，以及是否为零值，返回 (GetValue,IsZero)
	Get() (T, bool)

	// Or 只有当前为零值时，则取传入的值，创建新的Optional返回，原Optional不改变
	Or(T) Optional[T]
	// OrElse 只有当前为零值时，则返回传入的值
	OrElse(T) T
	// OrGet 只有当前为零值时，则执行传入的GetFun
	OrGet(GetFun[T]) T

	// Map 如果当前值不为零值，则映射到另外的Optional
	//Map(MapFun[T, any]) Optional[any]
	//MapZero(mapper MapFun[T, any]) Optional[any]
}

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
