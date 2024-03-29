package agent

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMemStat(t *testing.T) {
	m := getRunTimeStat()
	assert.NotNil(t, m, "getRunTimeStat should not return nil")

	keys := []string{"RandomValue", "Alloc", "BuckHashSys", "Frees", "GCCPUFraction", "GCSys", "HeapAlloc", "HeapIdle", "HeapInuse", "HeapObjects", "HeapReleased", "HeapSys", "LastGC", "Lookups", "MCacheInuse", "MCacheSys", "Mallocs", "NextGC", "NumForcedGC", "NumGC", "OtherSys", "PauseTotalNs", "StackInuse", "StackSys", "Sys", "TotalAlloc"}
	for _, key := range keys {
		_, ok := m[key]
		assert.True(t, ok, "Expected key %q to exist in the returned map", key)
	}
}
