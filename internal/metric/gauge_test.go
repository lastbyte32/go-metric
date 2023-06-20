package metric

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGauge_GetName(t *testing.T) {
	m := NewGauge("gauge1", 3.14)
	expectedName := "gauge1"
	assert.Equal(t, expectedName, m.GetName())
}

func TestGauge_GetType(t *testing.T) {
	m := NewGauge("gauge1", 3.14)
	assert.Equal(t, GAUGE, m.GetType())
}

func TestGauge_ToString(t *testing.T) {
	m := NewGauge("gauge1", 3.14)
	expectedString := "3.14"
	assert.Equal(t, expectedString, m.ToString())
}

func TestGauge_SetValue(t *testing.T) {
	m := NewGauge("gauge1", 3.14)
	err := m.SetValue("2.718")
	assert.NoError(t, err)
	expectedValue := 2.718
	assert.Equal(t, expectedValue, m.(*gauge).value)
}

func TestGauge_MarshalJSON(t *testing.T) {
	m := NewGauge("gauge1", 3.14)
	expectedJSON := `{"id":"gauge1","type":"gauge","value":3.14}`
	jsonData, err := m.MarshalJSON()
	assert.NoError(t, err)
	assert.Equal(t, []byte(expectedJSON), jsonData)
}

func TestGauge_SetHash(t *testing.T) {
	m := NewGauge("gauge1", 3.14)
	key := "mySecretKey"
	err := m.SetHash(key)
	assert.NoError(t, err)
	assert.NotEmpty(t, m.(*gauge).hash)
}
