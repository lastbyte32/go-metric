package handlers

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type GaugeMockStorage struct {
}

func (s *GaugeMockStorage) Update(name string, value float64) {

}

type CounterMockStorage struct {
}

func (s *CounterMockStorage) Update(name string, value int64) {

}

func TestUpdateHandle(t *testing.T) {
	gaugeStorage := &GaugeMockStorage{}
	countersStorage := &CounterMockStorage{}
	const (
		contentTypeOk = "text/plain; charset=utf-8"
		contentTypeNo = ""
	)

	type want struct {
		code        int
		contentType string
	}
	tests := []struct {
		name string
		url  string
		want want
	}{
		{
			name: "Counter err value",
			url:  "counter/test/none",
			want: want{
				code:        http.StatusBadRequest,
				contentType: contentTypeNo,
			},
		},
		{
			name: "Gauge err value",
			url:  "gauge/test/none",
			want: want{
				code:        http.StatusBadRequest,
				contentType: contentTypeNo,
			},
		},

		{
			name: "Counter OK",
			url:  "counter/test/1",
			want: want{
				code:        http.StatusOK,
				contentType: contentTypeOk,
			},
		},
		{
			name: "Gauge OK",
			url:  "gauge/test/1",
			want: want{
				code:        http.StatusOK,
				contentType: contentTypeOk,
			},
		},

		{
			name: "Metric type err",
			url:  "test/test/1",
			want: want{
				code:        http.StatusNotImplemented,
				contentType: contentTypeNo,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			target := "/update/" + tt.url
			request := httptest.NewRequest(http.MethodPost, target, nil)
			request.Header.Set("Content-Type", "text/plain")
			w := httptest.NewRecorder()
			h := UpdateHandle(gaugeStorage, countersStorage)
			h(w, request)
			res := w.Result()
			defer res.Body.Close()

			assert.Equal(t, tt.want.code, res.StatusCode, "incorrect resp status")
			assert.Equal(t, tt.want.contentType, res.Header.Get("Content-Type"), "incorrect header")

		})
	}
}
