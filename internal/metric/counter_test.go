package metric

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCounter_GetName(t *testing.T) {
	counterName := "counter1"
	m := NewCounter(counterName, 10)
	assert.Equal(t, counterName, m.GetName())
}

func TestCounter_GetType(t *testing.T) {
	m := NewCounter("counter1", 10)
	assert.Equal(t, COUNTER, m.GetType())
}

func TestCounter_ToString(t *testing.T) {
	m := NewCounter("counter1", 10)
	expectedString := "10"
	assert.Equal(t, expectedString, m.ToString())
}

func TestCounter_SetValue(t *testing.T) {
	m := NewCounter("counter1", 10)
	err := m.SetValue("5")
	assert.NoError(t, err)
	expectedValue := int64(15)
	assert.Equal(t, expectedValue, m.(*counter).value)
	err = m.SetValue("bad")
	assert.Error(t, err)
}

func TestCounter_MarshalJSON(t *testing.T) {
	m := NewCounter("counter1", 10)
	expectedJSON := `{"id":"counter1","type":"counter","delta":10}`
	jsonData, err := m.MarshalJSON()
	fmt.Println(string(jsonData))
	assert.NoError(t, err)
	assert.Equal(t, []byte(expectedJSON), jsonData)
}

func TestCounter_SetHash(t *testing.T) {
	m := NewCounter("counter1", 10)
	key := "mySecretKey"
	err := m.SetHash(key)
	assert.NoError(t, err)
	assert.NotEmpty(t, m.(*counter).hash)
}
