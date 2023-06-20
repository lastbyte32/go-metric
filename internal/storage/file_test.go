package storage

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestFileStorage_openFile(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	sugar := logger.Sugar()

	tmpDir, _ := os.MkdirTemp("", "storage_test")
	defer os.RemoveAll(tmpDir)

	store := &fileStorage{
		logger:     sugar,
		pathToFile: filepath.Join(tmpDir, "metrics.json"),
	}

	file, err := store.openFile(writeOrCreateMode)

	assert.NotNil(t, file)
	assert.NoError(t, err)

	file.Close()
}

func TestFileStorage_GetHash(t *testing.T) {
	storage := &fileStorage{}

	data := []byte{}
	expectedHash := "d41d8cd98f00b204e9800998ecf8427e"
	hash := storage.getHash(data)
	assert.Equal(t, expectedHash, hash, "Empty data hash does not match")

	data = []byte("Hello, world!")
	expectedHash = "6cd3556deb0da54bca060b4c39479839"
	hash = storage.getHash(data)
	assert.Equal(t, expectedHash, hash, "Non-empty data hash does not match")
}

func TestFileStorage_HasChanges(t *testing.T) {
	storage := &fileStorage{hash: "test_hash"}

	data := []byte("test_data")
	hasChanges := storage.hasChanges(data)
	assert.True(t, hasChanges, "Expected changes")

	// Test case 2: Different hash than previous
	data = []byte("test_data")
	hasChanges = storage.hasChanges(data)
	assert.False(t, hasChanges, "Expected no changes")

	data = []byte{}
	hasChanges = storage.hasChanges(data)
	assert.True(t, hasChanges, "Expected changes (empty data)")
}
