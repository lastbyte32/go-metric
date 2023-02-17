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

type MockStorage struct {
	mock.Mock
}

func (m *MockStorage) Get(name string) (storage.Metric, bool) {
	rs := m.Called(name)
	return rs.Get(0).(storage.Metric), rs.Bool(1)
}

func (m *MockStorage) All() map[string]storage.Metric {
	rs := m.Called()
	return rs.Get(0).(map[string]storage.Metric)
}

func (m *MockStorage) Update(string, storage.Metric) {}

func TestGetMetric(t *testing.T) {

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
		name    string
		req     *chi.Context
		storage storage.Storage
		want    want
	}{

		{
			name: "Counter get value OK",
			req: func() *chi.Context {
				rctx := chi.NewRouteContext()
				rctx.URLParams.Add("type", string(storage.COUNTER))
				rctx.URLParams.Add("name", "my_counter")
				return rctx
			}(),
			storage: func() storage.Storage {
				r := new(MockStorage)
				r.On("Get", "my_counter").
					Return(storage.NewMetric(
						"test",
						storage.COUNTER,
						0,
						100,
					),
						true)
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
				rctx.URLParams.Add("type", string(storage.GAUGE))
				rctx.URLParams.Add("name", "my_gauge")
				return rctx
			}(),
			storage: func() storage.Storage {
				r := new(MockStorage)
				r.On("Get", "my_gauge").
					Return(storage.NewMetric(
						"my_gauge",
						storage.GAUGE,
						1.1,
						0,
					),
						true)
				return r
			}(),

			want: want{
				code:        http.StatusOK,
				contentType: contentTypeOk,
				result:      "1.100",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			handler := NewHandler(tt.storage)

			request := httptest.NewRequest(http.MethodGet, "/", nil)
			request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, tt.req))

			w := httptest.NewRecorder()
			h := http.HandlerFunc(handler.GetMetric)
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
