package handlers

import (
	"fmt"

	"go.uber.org/zap"

	"github.com/lastbyte32/go-metric/internal/config"
	"github.com/lastbyte32/go-metric/internal/storage"
	"github.com/lastbyte32/go-metric/pkg/utils/crypto"
)

type decrypter interface {
	Decrypt(encryptedMessage []byte) ([]byte, error)
}

type handler struct {
	metricsStorage storage.IStorage
	logger         *zap.SugaredLogger
	config         config.Configurator
	isDecrypt      bool
	decrypter      decrypter
}

func NewHandler(s storage.IStorage, l *zap.SugaredLogger, c config.Configurator) (*handler, error) {
	h := &handler{
		metricsStorage: s,
		logger:         l,
		config:         c,
		isDecrypt:      false,
	}
	keyPath := c.GetCryptoKeyPath()
	if keyPath != "" {
		decrypter, err := crypto.NewEncryptor(keyPath)
		if err != nil {
			return nil, fmt.Errorf("error on create decryptor: %v", err)
		}
		h.decrypter = decrypter
		h.isDecrypt = true
	}

	return h, nil
}
