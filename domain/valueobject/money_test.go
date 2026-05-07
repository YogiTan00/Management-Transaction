package valueobject_test

import (
	"github.com/YogiTan00/Management-Transaction/domain/valueobject"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMoney(t *testing.T) {
	t.Run("set 0 if money is minus", func(t *testing.T) {
		result := valueobject.NewMoney(int64(-1))
		assert.Equal(t, valueobject.Money(0), result)
	})
	t.Run("set value if correct logic", func(t *testing.T) {
		result := valueobject.NewMoney(int64(10000))
		assert.Equal(t, valueobject.Money(10000), result)
	})
}
