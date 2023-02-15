package handlers

import (
	"fmt"
	"github.com/lastbyte32/go-metric/internal/server/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndex(t *testing.T) {
	h := &Main{
		GaugeStorage: func() storage.GaugeStorage {
			r := new(GaugeMockStorage)
			r.On("Get", "my_gauge").Return(float64(10000.1), true)
			r.On("All").Return(map[string]float64{
				"my_gauge": 10000.1,
			})

			return r
		}(),
		CountersStorage: func() storage.CounterStorage {
			r := new(CounterMockStorage)
			r.On("Get", "my_counter").Return(int64(100), true)
			r.On("All").Return(map[string]int64{
				"my_counter": 100,
			})

			return r
		}(),
	}

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	handler := http.HandlerFunc(h.Index)

	handler.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code, fmt.Sprintf("handler returned wrong status code: got %v want %v", w.Code, http.StatusOK))
	require.Contains(t, w.Body.String(), "<title>Metrics</title>", "returned wrong body")

	values := map[string]string{
		"my_gauge":   "10000.100000",
		"my_counter": "100",
	}
	for k, v := range values {
		assert.Contains(t, w.Body.String(), fmt.Sprintf("<li><b>%s:</b> %s</li>", k, v), fmt.Sprintf("not fount %s = %s", k, v))
	}

}

func TestIndexEmpty(t *testing.T) {
	h := &Main{
		GaugeStorage: func() storage.GaugeStorage {
			r := new(GaugeMockStorage)
			r.On("All").Return(map[string]float64{})

			return r
		}(),
		CountersStorage: func() storage.CounterStorage {
			r := new(CounterMockStorage)
			r.On("All").Return(map[string]int64{})
			return r
		}(),
	}

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	handler := http.HandlerFunc(h.Index)
	handler.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code, fmt.Sprintf("handler returned wrong status code: got %v want %v", w.Code, http.StatusOK))
	require.Contains(t, w.Body.String(), "<li><b>No metrics</b></li>", "returned wrong body")

}
