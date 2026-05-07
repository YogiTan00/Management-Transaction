package valueobject_test

import (
	"github.com/YogiTan00/Management-Transaction/domain/valueobject"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewDiscount(t *testing.T) {
	t.Run("set 0 if value is minus", func(t *testing.T) {
		result := valueobject.NewDiscount(-1)
		assert.Equal(t, valueobject.Discount(0), result)
	})
	t.Run("set value if correct logic", func(t *testing.T) {
		result := valueobject.NewDiscount(10)
		assert.Equal(t, valueobject.Discount(10), result)
	})
	t.Run("set 100 if value more than 100", func(t *testing.T) {
		result := valueobject.NewDiscount(100000)
		assert.Equal(t, valueobject.Discount(100), result)
	})
}
