package metric

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGauge(t *testing.T) {
	name := "test_gauge"
	value := float64(10.100)
	gauge := NewGauge(name, value)

	t.Run("GetName", func(t *testing.T) {
		assert.Equal(t, name, gauge.GetName())
	})

	t.Run("GetType", func(t *testing.T) {
		assert.Equal(t, GAUGE, gauge.GetType())
	})

	t.Run("ToString", func(t *testing.T) {
		assert.Equal(t, "10.100", gauge.ToString())
	})

	t.Run("SetValue", func(t *testing.T) {
		err := gauge.SetValue("5.555")
		assert.NoError(t, err)
		assert.Equal(t, "5.555", gauge.ToString())

		err = gauge.SetValue("invalid_value")
		assert.Error(t, err)
	})
}
