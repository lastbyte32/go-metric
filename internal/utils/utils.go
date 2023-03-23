package utils

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"
)

func StringToFloat64(value string) (float64, error) {
	f, err := strconv.ParseFloat(value, 64)
	if err != nil {
		fmt.Println("StringToFloat64 err parse")
		return float64(0), err
	}
	return f, nil
}

func StringToInt64(value string) (int64, error) {
	f, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		fmt.Println("StringToInt64 err parse")
		return int64(0), err
	}
	return f, nil
}

func FloatToString(value float64) string {
	return strconv.FormatFloat(value, 'f', -1, 64)
}

func GetMD5Sum(data []byte) []byte {
	hash := md5.New()
	hash.Write(data)
	return hash.Sum(nil)
}

func GetMD5Hash(data []byte) string {
	return hex.EncodeToString(GetMD5Sum(data))
}

func GetSha256Hash(part, key string) (string, error) {
	if key == "" {
		return "", errors.New("empty key")
	}

	h := hmac.New(sha256.New, []byte(key))

	if _, err := h.Write([]byte(part)); err != nil {
		return "", err
	}

	return hex.EncodeToString(h.Sum(nil)), nil
}
