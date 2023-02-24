package metric

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewByString_Gauge(t *testing.T) {
	name := "test_gauge"
	value := "123.450"
	metricType := GAUGE

	metric, err := NewByString(name, value, metricType)

	assert.NoError(t, err)
	assert.NotNil(t, metric)
	assert.Equal(t, name, metric.GetName())
	assert.Equal(t, GAUGE, metric.GetType())
	assert.Equal(t, value, metric.ToString())

}

func TestNewByString_Counter(t *testing.T) {
	name := "test_counter"
	value := "123"
	metricType := COUNTER

	metric, err := NewByString(name, value, metricType)

	assert.NoError(t, err)
	assert.NotNil(t, metric)
	assert.Equal(t, name, metric.GetName())
	assert.Equal(t, COUNTER, metric.GetType())
	assert.Equal(t, value, metric.ToString())
}

func TestNewByString_WrongMetricType(t *testing.T) {
	name := "test_wrong_type"
	value := "123.45"
	metricType := "WRONG_TYPE"

	metric, err := NewByString(name, value, MType(metricType))

	assert.Error(t, err)
	assert.Nil(t, metric)
	assert.Equal(t, errors.New("NewByString: wrong metric type"), err)
}

func TestNewByString_InvalidValueForGauge(t *testing.T) {
	name := "test_invalid_gauge"
	value := "invalid_value"
	metricType := GAUGE

	metric, err := NewByString(name, value, metricType)

	assert.Error(t, err)
	assert.Nil(t, metric)
}

func TestNewByString_InvalidValueForCounter(t *testing.T) {
	name := "test_invalid_counter"
	value := "invalid_value"
	metricType := COUNTER

	metric, err := NewByString(name, value, metricType)

	assert.Error(t, err)
	assert.Nil(t, metric)
}
