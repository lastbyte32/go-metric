package agent

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func Test_poolGauge(t *testing.T) {
	gauges := poolGauge()

	require.NotEmpty(t, gauges, "Returned slice is empty")
	assert.Equal(t, gauges[0].name, "RandomValue")

	expectedGauges := []string{
		"Alloc",
		"BuckHashSys",
		"Frees",
		"GCCPUFraction",
		"GCSys",
		"HeapAlloc",
		"HeapIdle",
		"HeapInuse",
		"HeapObjects",
		"HeapReleased",
		"HeapSys",
		"LastGC",
		"Lookups",
		"MCacheInuse",
		"MCacheSys",
		"Mallocs",
		"NextGC",
		"NumForcedGC",
		"NumGC",
		"OtherSys",
		"PauseTotalNs",
		"RandomValue",
		"StackInuse",
		"StackSys",
		"Sys",
		"TotalAlloc",
	}
	for _, expected := range expectedGauges {
		found := false
		for _, g := range gauges {
			if g.name == expected {
				found = true
				assert.IsType(t, g.value, float64(0), "wrong type")
				break
			}
		}
		assert.True(t, found, "Expected gauge %s not found", expected)
	}
}

func TestPoolCounter(t *testing.T) {
	counters := poolCounter(10)
	require.NotEmpty(t, counters, "Unexpected output length")
	assert.Equal(t, int64(10), counters[0].value, "Unexpected counter value")
}

func TestCounterToString(t *testing.T) {
	c := counter{"TestCounter", 123}
	assert.Equal(t, "123", c.toString(), "Unexpected string representation of counter")
}

func TestCounterGetURLUpdateParam(t *testing.T) {
	c := counter{"TestCounter", 123}
	assert.Equal(t, "counter/TestCounter/123", c.getURLUpdateParam(), "Unexpected URL parameter")
}

func TestSendReportSuccess(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/update/counter/my_counter/42", r.URL.Path)
		assert.Equal(t, "text/plain", r.Header.Get("Content-Type"))
		body, err := io.ReadAll(r.Body)
		assert.NoError(t, err)
		assert.Equal(t, "42", string(body))
		w.WriteHeader(http.StatusOK)
	}))

	defer server.Close()

	c := counter{name: "my_counter", value: 42}

	err := c.sendReport(server.URL[7:], time.Second)
	assert.NoError(t, err)
}
