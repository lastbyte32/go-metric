package metric

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewByString(t *testing.T) {
	type args struct {
		name       string
		value      string
		metricType MType
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test_gauge",
			args: args{
				name:       "test_gauge",
				value:      "10.5",
				metricType: GAUGE,
			},
			wantErr: false,
		},

		{
			name: "test_gauge",
			args: args{
				name:       "test_gauge",
				value:      "wrong",
				metricType: GAUGE,
			},
			wantErr: true,
		},

		{
			name: "test_counter",
			args: args{
				name:       "test_counter",
				value:      "10",
				metricType: COUNTER,
			},
			wantErr: false,
		},

		{
			name: "test_Counter2",
			args: args{
				name:       "test_counter",
				value:      "wrong",
				metricType: COUNTER,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewByString(tt.args.name, tt.args.value, tt.args.metricType)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			assert.Equal(t, tt.name, got.GetName())
			assert.Equal(t, tt.args.metricType, got.GetType())
			assert.Equal(t, tt.args.value, got.ToString())
			assert.ObjectsAreEqual(gauge{}, got)

		})
	}
}

func TestCounterGetValueAsString(t *testing.T) {
	m := Metrics{
		MType: string(COUNTER),
		Delta: new(int64),
	}
	*m.Delta = 10
	expectedCounter := "10"
	assert.Equal(t, expectedCounter, m.GetValueAsString())
}
func TestGaugeGetValueAsString(t *testing.T) {
	m := Metrics{
		MType: string(GAUGE),
		Value: new(float64),
	}
	value := 25.5
	expectedGauge := "25.5"
	m.Value = &value
	assert.Equal(t, expectedGauge, m.GetValueAsString())
}

func TestGetStringToSign(t *testing.T) {
	m := Metrics{
		ID:    "metric1",
		MType: "counter",
		Delta: new(int64),
		Value: nil,
	}

	*m.Delta = 10

	expectedCounter := "metric1:counter:10"
	assert.Equal(t, expectedCounter, m.GetStringToSign())

	m.MType = "gauge"
	value := 25.5
	m.Value = &value

	expectedGauge := "metric1:gauge:25.500000"
	assert.Equal(t, expectedGauge, m.GetStringToSign())

	m.MType = "err_type"
	expectedDefault := ""
	assert.Equal(t, expectedDefault, m.GetStringToSign())
}
