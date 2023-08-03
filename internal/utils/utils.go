package utils

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"net"
	"strconv"
)

// StringToFloat64 - функция преобразования строки в число с плавающей точкой.
func StringToFloat64(value string) (float64, error) {
	f, err := strconv.ParseFloat(value, 64)
	if err != nil {
		fmt.Println("StringToFloat64 err parse")
		return float64(0), err
	}
	return f, nil
}

// StringToInt64 - функция преобразования строки в число.
func StringToInt64(value string) (int64, error) {
	f, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		fmt.Println("StringToInt64 err parse")
		return int64(0), err
	}
	return f, nil
}

// FloatToString - функция преобразования числа в строку.
func FloatToString(value float64) string {
	return strconv.FormatFloat(value, 'f', -1, 64)
}

// GetMD5Sum - функция получения MD5-хеша в виде слайса байт.
func GetMD5Sum(data []byte) []byte {
	hash := md5.New() //nolint:gosec
	hash.Write(data)
	return hash.Sum(nil)
}

// GetMD5Hash - функция получения MD5-хеша в виде строки.
func GetMD5Hash(data []byte) string {
	return hex.EncodeToString(GetMD5Sum(data))
}

// GetSha256Hash - функция получения SHA-256-хеша.
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

func GetFirstHostIPv4Addr() (net.IP, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}
	for _, addr := range addrs {
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			return ipNet.IP, nil
		}
	}

	return net.IP{}, fmt.Errorf("can't get IPv4 address")
}
