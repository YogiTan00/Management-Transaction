package valueobject_test

import (
	"github.com/YogiTan00/Management-Transaction/domain/valueobject"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTypeSize(t *testing.T) {
	t.Run("set value if value 0", func(t *testing.T) {
		result := valueobject.NewTypeSize(0)
		assert.Equal(t, valueobject.TypeSize(0), result)
	})
	t.Run("set value if value 1", func(t *testing.T) {
		result := valueobject.NewTypeSize(1)
		assert.Equal(t, valueobject.TypeSize(1), result)
	})
	t.Run("set value to default if not whitelisted", func(t *testing.T) {
		result := valueobject.NewTypeSize(99)
		assert.Equal(t, valueobject.TypeSize(0), result)
	})
}
