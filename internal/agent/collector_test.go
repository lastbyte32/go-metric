package agent

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_poolGauge(t *testing.T) {
	var testGauge []gauge
	poolGauge(&testGauge)

	require.NotEmpty(t, testGauge, "Expected non-empty gauge list, but got empty")
	// Проверяем, что каждый показатель имеет правильный формат
	//for _, g := range testGauge {
	//}
}

func TestPoolGaugeWithNilPointer(t *testing.T) {
	var g *[]gauge

	assert.Panics(t, func() {
		poolGauge(g)
	}, "Expected panic, but it did not happen")
}

func TestPoolCounter(t *testing.T) {
	var counters []counter
	cnt := int64(5)

	poolCounter(&counters, cnt)

	assert.Len(t, counters, 1, "Unexpected number of counters")
}

func TestPoolCounterWithNilPointer(t *testing.T) {
	var counters *[]counter
	cnt := int64(5)

	assert.Panics(t, func() {
		poolCounter(counters, cnt)
	}, "Expected panic, but it did not happen")
}
