package middleware

import (
	"compress/gzip"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompressor(t *testing.T) {
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	request.Header.Set("Accept-Encoding", "gzip")

	recorder := httptest.NewRecorder()

	compressedHandler := Compressor(testHandler)
	compressedHandler.ServeHTTP(recorder, request)
	response := recorder.Result()
	response.Body.Close()
	assert.Equal(t, "gzip", response.Header.Get("Content-Encoding"))

	gzipReader, err := gzip.NewReader(response.Body)
	assert.NoError(t, err)
	defer gzipReader.Close()

	body, err := io.ReadAll(gzipReader)
	assert.NoError(t, err)

	assert.Equal(t, "Hello, world!", string(body))
}
