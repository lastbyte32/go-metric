package storage

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGaugeMemStorage_All(t *testing.T) {
	g := NewGaugeMemStorage()
	g.Update("foo_gauge", 99.99)
	g.Update("bar_gauge", 100.001)
	all := g.All()
	assert.Equal(t, 2, len(all), fmt.Sprintf("Expected 2 gauges, but got %d", len(all)))

	assert.Equal(t, 99.99, all["foo_gauge"], fmt.Sprintf("Expected foo_gauge to be 1, but got %f", all["foo_gauge"]))
	assert.Equal(t, 100.001, all["bar_gauge"], fmt.Sprintf("Expected bar_gauge to be 1, but got %f", all["bar_gauge"]))
}

func TestGaugeMemStorage_Get(t *testing.T) {
	s := NewGaugeMemStorage()
	s.Update("foo_value", 1.1)

	value, exist := s.Get("foo_value")
	assert.True(t, exist, "Failed to get value for foo_value")
	assert.Equal(t, 1.1, value, fmt.Sprintf("Value for foo_value should be 1 but got %f", value))

	value, exist = s.Get("non-existent")
	assert.False(t, exist)
	assert.Equal(t, float64(0), value)
}

func TestGaugeMemStorage_Update(t *testing.T) {

	type fields struct {
		values map[string]float64
	}

	type args struct {
		name  string
		value float64
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
	}{
		{
			"Test1 update and get",
			fields{map[string]float64{"test": 99.99}},
			args{"test", 3},
			3,
		},

		{
			"Test2 override value by key",
			fields{map[string]float64{"test": 100.99}},
			args{"test2", 99.99},
			99.99,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GaugeMemStorage{
				values: tt.fields.values,
			}
			g.Update(tt.args.name, tt.args.value)

			value, exist := g.Get(tt.args.name)
			assert.True(t, exist)
			assert.Equal(t, tt.want, value)

		})
	}
}

func TestNewGaugeMemStorage(t *testing.T) {
	tests := []struct {
		name string
		want GaugeStorage
	}{
		{"First",
			NewGaugeMemStorage(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NotNil(t, tt.want, "NewGaugeMemStorage() returned nil")
			assert.Equalf(t, tt.want, NewGaugeMemStorage(), "NewGaugeMemStorage()")
		})
	}
}
