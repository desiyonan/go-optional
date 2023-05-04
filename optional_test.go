package optional

import (
	. "github.com/stretchr/testify/assert"
	"testing"
)

func Test_Of(t *testing.T) {
	intOpt := Of(1)
	var intZero int
	intZeroOpt := Of(intZero)

	Equal(t, false, intOpt.IsZero())
	Equal(t, 1, intOpt.GetValue())
	Equal(t, 1, intOpt.OrElse(2))
	Equal(t, true, intZeroOpt.IsZero())
	Equal(t, 0, intZeroOpt.GetValue())
	Equal(t, 2, intZeroOpt.OrElse(2))
}

func Test_Any(t *testing.T) {
	Equal(t, 0, Any(0).GetValue())
	Equal(t, 1, Any(1).GetValue())

	Equal(t, 1, Any(0, 1, 2, 3).GetValue())
	Equal(t, 1, Any(1, 2, 3).GetValue())

	Equal(t, 1, Any(0, 0, 0, 1).GetValue())
	Equal(t, 1, Any(1, 2, 0, 0, 3).GetValue())
}
