package storage

import (
	"io/ioutil"
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

	tmpDir, _ := ioutil.TempDir("", "storage_test")
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
