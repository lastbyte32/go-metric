package agent

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestTransmitPlainText_Success(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer mockServer.Close()

	err := transmitPlainText(mockServer.URL, "test", 5*time.Second)

	assert.NoError(t, err)
}

func TestTransmitPlainText_Fail(t *testing.T) {
	assert.Error(t, transmitPlainText("invalid", "test", 5*time.Second))
}
