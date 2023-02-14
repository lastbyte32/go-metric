package storage

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCounterMemStorage_All(t *testing.T) {
	s := NewCounterMemStorage()
	s.Update("foo_counter", 1)
	s.Update("bar_counter", 2)
	all := s.All()
	assert.Equal(t, 2, len(all), fmt.Sprintf("Expected 2 counters, but got %d", len(all)))
	assert.Equal(t, int64(1), all["foo_counter"], fmt.Sprintf("Expected foo_counter to be 1, but got %d", all["foo_counter"]))
	assert.Equal(t, int64(2), all["bar_counter"], fmt.Sprintf("Expected bar_counter to be 1, but got %d", all["bar_counter"]))

}

func TestCounterMemStorage_Get(t *testing.T) {
	s := NewCounterMemStorage()
	s.Update("foo_counter", 1)

	value, exist := s.Get("foo_counter")
	assert.True(t, exist, "Failed to get value for foo_counter")
	assert.Equal(t, int64(1), value, fmt.Sprintf("Value for test_counter should be 1 but got %d", value))

	value, exist = s.Get("non-existent")
	assert.False(t, exist)
	assert.Equal(t, int64(0), value)
}

func TestCounterMemStorage_Update(t *testing.T) {
	s := NewCounterMemStorage()

	s.Update("foo", 1)
	value, exist := s.Get("foo")
	assert.True(t, exist)
	assert.Equal(t, int64(1), value)

	s.Update("foo", 2)
	value, exist = s.Get("foo")
	assert.True(t, exist)
	assert.Equal(t, int64(3), value)
}

func TestNewCounterMemStorage(t *testing.T) {
	s := NewCounterMemStorage()
	assert.NotNil(t, s, "NewCounterMemStorage() returned nil")
}
