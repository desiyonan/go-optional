package optional

import (
	. "github.com/stretchr/testify/assert"
	"testing"
)

type testObj struct {
	name string
}

func Test_Some(t *testing.T) {
	opt := Some(0)

	Equal(t, true, opt.IsZero())
	Equal(t, 0, opt.GetValue())

	Equal(t, 2, opt.Or(2).GetValue())
	Equal(t, true, opt.IsZero())

	Equal(t, 2, opt.OrElse(2))
	Equal(t, true, opt.IsZero())

	Equal(t, 3, opt.OrGet(func() int { return 3 }))
	Equal(t, true, opt.IsZero())

	opt2 := Some("")

	Equal(t, true, opt2.IsZero())
	Equal(t, "", opt2.GetValue())

	Equal(t, "a", opt2.Or("a").GetValue())
	Equal(t, true, opt2.IsZero())

	Equal(t, "b", opt2.OrElse("b"))
	Equal(t, true, opt2.IsZero())

	Equal(t, "c", opt2.OrGet(func() string { return "c" }))
	Equal(t, true, opt2.IsZero())

	obj := testObj{}

	opt3 := Some(obj)

	Equal(t, true, opt3.IsZero())
	Equal(t, obj, opt3.GetValue())

	Equal(t, "a", opt3.Or(testObj{"a"}).GetValue().name)
	Equal(t, false, opt3.Or(testObj{"a"}).IsZero())
	Equal(t, true, opt3.IsZero())

	Equal(t, "b", opt3.OrElse(testObj{"b"}).name)
	Equal(t, true, opt3.IsZero())

	Equal(t, "c", opt3.OrGet(func() testObj { return testObj{"c"} }).name)
	Equal(t, true, opt3.IsZero())

	var ptr *testObj

	opt4 := Some(ptr)

	Equal(t, true, opt4.IsZero())
	Equal(t, ptr, opt4.GetValue())

	Equal(t, "a", opt4.Or(&testObj{"a"}).GetValue().name)
	Equal(t, false, opt4.Or(&testObj{"a"}).IsZero())
	Equal(t, true, opt4.IsZero())

	Equal(t, "b", opt4.OrElse(&testObj{"b"}).name)
	Equal(t, true, opt4.IsZero())

	Equal(t, "c", opt4.OrGet(func() *testObj { return &testObj{"c"} }).name)
	Equal(t, true, opt4.IsZero())
}
