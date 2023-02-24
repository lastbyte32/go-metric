package metric

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCounter(t *testing.T) {
	name := "test_counter"
	value := int64(10)
	counter := NewCounter(name, value)

	t.Run("GetName", func(t *testing.T) {
		assert.Equal(t, name, counter.GetName())
	})

	t.Run("GetType", func(t *testing.T) {
		assert.Equal(t, COUNTER, counter.GetType())
	})

	t.Run("ToString", func(t *testing.T) {
		assert.Equal(t, "10", counter.ToString())
	})

	t.Run("SetValue", func(t *testing.T) {
		err := counter.SetValue("5")
		assert.NoError(t, err)
		assert.Equal(t, "15", counter.ToString())

		err = counter.SetValue("invalid_value")
		assert.Error(t, err)
	})
}
