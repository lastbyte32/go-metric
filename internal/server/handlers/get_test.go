package handlers

import (
	"context"
	"github.com/go-chi/chi"
	"github.com/lastbyte32/go-metric/internal/server/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type GaugeMockStorage struct {
	mock.Mock
}

func (s *GaugeMockStorage) Get(name string) (float64, bool) {
	rs := s.Called(name)
	return rs.Get(0).(float64), rs.Bool(1)
}

func (s *GaugeMockStorage) All() map[string]float64 {
	rs := s.Called()
	return rs.Get(0).(map[string]float64)
}

func (s *GaugeMockStorage) Update(name string, value float64) {

}

type CounterMockStorage struct {
	mock.Mock
}

func (s *CounterMockStorage) Get(name string) (int64, bool) {
	rs := s.Called(name)
	return rs.Get(0).(int64), rs.Bool(1)
}

func (s *CounterMockStorage) All() map[string]int64 {
	rs := s.Called()
	return rs.Get(0).(map[string]int64)
}

func (s *CounterMockStorage) Update(name string, value int64) {

}

func TestGetHandle(t *testing.T) {

	const (
		contentTypeOk = "text/plain"
		contentTypeNo = ""
	)

	type want struct {
		code        int
		contentType string
		result      string
	}
	tests := []struct {
		name            string
		req             *chi.Context
		gaugeStorage    storage.GaugeStorage
		countersStorage storage.CounterStorage
		want            want
	}{
		{
			name: "Counter get value OK",
			req: func() *chi.Context {
				rctx := chi.NewRouteContext()
				rctx.URLParams.Add("type", "counter")
				rctx.URLParams.Add("name", "my_counter")
				return rctx
			}(),
			gaugeStorage: new(GaugeMockStorage),

			countersStorage: func() storage.CounterStorage {
				r := new(CounterMockStorage)
				r.On("Get", "my_counter").Return(int64(100), true)
				return r
			}(),

			want: want{
				code:        http.StatusOK,
				contentType: contentTypeOk,
				result:      "100",
			},
		},
		{
			name: "Gauge get value OK",
			req: func() *chi.Context {
				rctx := chi.NewRouteContext()
				rctx.URLParams.Add("type", "gauge")
				rctx.URLParams.Add("name", "my_gauge")
				return rctx
			}(),
			gaugeStorage: func() storage.GaugeStorage {
				r := new(GaugeMockStorage)
				r.On("Get", "my_gauge").Return(float64(10000.1), true)
				return r
			}(),

			countersStorage: new(CounterMockStorage),

			want: want{
				code:        http.StatusOK,
				contentType: contentTypeOk,
				result:      "10000.1",
			},
		},

		//{
		//	name: "Gauge err value",
		//	url:  "gauge/test/none",
		//	want: want{
		//		code:        http.StatusBadRequest,
		//		contentType: contentTypeNo,
		//	},
		//},
		//
		//{
		//	name: "Counter OK",
		//	url:  "counter/test/1",
		//	want: want{
		//		code:        http.StatusOK,
		//		contentType: contentTypeOk,
		//	},
		//},
		//{
		//	name: "Gauge OK",
		//	url:  "gauge/test/1",
		//	want: want{
		//		code:        http.StatusOK,
		//		contentType: contentTypeOk,
		//	},
		//},
		//
		//{
		//	name: "Metric type err",
		//	url:  "test/test/1",
		//	want: want{
		//		code:        http.StatusNotImplemented,
		//		contentType: contentTypeNo,
		//	},
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			handler := &Main{
				GaugeStorage:    tt.gaugeStorage,
				CountersStorage: tt.countersStorage,
			}

			request := httptest.NewRequest(http.MethodGet, "/", nil)
			request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, tt.req))

			w := httptest.NewRecorder()
			h := http.HandlerFunc(handler.GetOneMetric)
			h.ServeHTTP(w, request)
			res := w.Result()
			res.Body.Close()

			assert.Equal(t, tt.want.code, res.StatusCode, "incorrect status")
			assert.Equal(t, tt.want.contentType, res.Header.Get("Content-Type"), "incorrect header")

			body, err := io.ReadAll(res.Body)
			require.NoError(t, err)
			assert.Equal(t, tt.want.result, string(body), "incorrect result")

		})
	}
}
