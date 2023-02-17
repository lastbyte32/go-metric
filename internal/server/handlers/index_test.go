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

	h := NewHandler(func() storage.Storage {
		r := new(MockStorage)
		r.On("Get", "my_gauge").
			Return(storage.NewMetric(
				"my_gauge",
				storage.GAUGE,
				10000.1,
				0,
			),
				true)

		r.On("Get", "my_gauge").
			Return(storage.NewMetric(
				"my_gauge",
				storage.GAUGE,
				10000.1,
				0,
			),
				true)

		r.On("All").Return(map[string]storage.Metric{
			"my_gauge": storage.NewMetric(
				"my_gauge",
				storage.GAUGE,
				10000.1,
				0,
			),
			"my_counter": storage.NewMetric(
				"my_counter",
				storage.COUNTER,
				0,
				100,
			),
		})

		return r
	}())

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
		"my_gauge":   "10000.100",
		"my_counter": "100",
	}
	for k, v := range values {
		assert.Contains(t, w.Body.String(), fmt.Sprintf("<li><b>%s:</b> %s</li>", k, v), fmt.Sprintf("not fount %s = %s", k, v))
	}

}

func TestIndexEmpty(t *testing.T) {
	h := NewHandler(func() storage.Storage {
		r := new(MockStorage)
		r.On("All").Return(map[string]storage.Metric{})
		return r
	}())

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
