package storage

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	mproto "github.com/lastbyte32/go-metric/api/proto"
	"github.com/lastbyte32/go-metric/internal/metric"
)

type mockMetric struct {
	name      string
	valueType metric.MType
	value     string
}

func (m *mockMetric) MarshalProtobuf() *mproto.Metric {
	//TODO implement me
	panic("implement me")
}

func (m *mockMetric) MarshalJSON() ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (m *mockMetric) SetHash(key string) error {
	//TODO implement me
	panic("implement me")
}

func (m *mockMetric) GetName() string {
	return m.name
}

func (m *mockMetric) GetType() metric.MType {
	return m.valueType
}

func (m *mockMetric) SetValue(value string) error {
	m.value = value
	return nil
}

func (m *mockMetric) ToString() string {
	return m.value
}

func TestMemoryStorage_Close(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	sugar := logger.Sugar()

	store := &memoryStorage{
		values: make(map[string]metric.IMetric),
		logger: sugar,
		mutex:  sync.RWMutex{},
	}

	err := store.Close()

	assert.NoError(t, err)
}

func TestMemoryStorage_Get(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	sugar := logger.Sugar()

	store := &memoryStorage{
		values: map[string]metric.IMetric{
			"metric1": &mockMetric{},
			"metric2": &mockMetric{},
		},
		logger: sugar,
		mutex:  sync.RWMutex{},
	}

	metric, found := store.Get("metric1")

	assert.True(t, found)
	assert.NotNil(t, metric)

	metric, found = store.Get("nonexistent")

	assert.False(t, found)
	assert.Nil(t, metric)
}

func TestMemoryStorage_Update(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	sugar := logger.Sugar()

	store := &memoryStorage{
		values: make(map[string]metric.IMetric),
		logger: sugar,
		mutex:  sync.RWMutex{},
	}

	err := store.Update("metric1", "1000", metric.COUNTER)

	assert.NoError(t, err)
	assert.Len(t, store.values, 1)
	assert.Contains(t, store.values, "metric1")

	err = store.Update("metric1", "10000", metric.COUNTER)

	assert.NoError(t, err)
	assert.Len(t, store.values, 1)
	assert.Contains(t, store.values, "metric1")

	err = store.Update("metric2", "10000", metric.GAUGE)

	assert.NoError(t, err)
	assert.Len(t, store.values, 2)
	assert.Contains(t, store.values, "metric1")
	assert.Contains(t, store.values, "metric2")
}

func TestMemoryStorage_All(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	sugar := logger.Sugar()

	store := &memoryStorage{
		values: map[string]metric.IMetric{
			"metric1": &mockMetric{},
			"metric2": &mockMetric{},
		},
		logger: sugar,
		mutex:  sync.RWMutex{},
	}

	allMetrics := store.All()

	assert.Len(t, allMetrics, 2)
	assert.Contains(t, allMetrics, "metric1")
	assert.Contains(t, allMetrics, "metric2")
}

func TestNewMemoryStorage(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	sugar := logger.Sugar()

	store := NewMemoryStorage(sugar)

	assert.NotNil(t, store)
}
