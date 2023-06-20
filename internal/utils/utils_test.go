package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStringToFloat64(t *testing.T) {

	tests := []struct {
		name    string
		args    string
		want    float64
		wantErr bool
	}{
		{
			"test0",
			"999.999",
			999.999,
			false,
		},

		{
			"test1",
			"-999.999",
			-999.999,
			false,
		},

		{
			"test2",
			"test",
			0,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StringToFloat64(tt.args)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			assert.Equalf(t, tt.want, got, "StringToFloat64() got = %v, want %v", got, tt.want)
		})
	}
}

func TestStringToInt64(t *testing.T) {
	tests := []struct {
		name    string
		args    string
		want    int64
		wantErr bool
	}{
		{
			name:    "test0",
			args:    "999",
			want:    999,
			wantErr: false,
		},
		{
			name:    "test1",
			args:    "-999",
			want:    -999,
			wantErr: false,
		},

		{
			name:    "test2",
			args:    "test",
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StringToInt64(tt.args)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			assert.Equalf(t, tt.want, got, "StringToInt64(%v)", tt.args)
		})
	}
}

func TestFloatToString(t *testing.T) {

	tests := []struct {
		name string
		args float64
		want string
	}{
		{
			"test0",
			1.1,
			"1.1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, FloatToString(tt.args), "FloatToString(%v)", tt.args)
		})
	}
}

func TestGetMD5Sum(t *testing.T) {
	data := []byte("test")
	expectedHash := []byte{0x9, 0x8f, 0x6b, 0xcd, 0x46, 0x21, 0xd3, 0x73, 0xca, 0xde, 0x4e, 0x83, 0x26, 0x27, 0xb4, 0xf6}

	assert.Equal(t, expectedHash, GetMD5Sum(data))
}

func TestGetMD5Hash(t *testing.T) {
	data := []byte("test")
	expectedHash := "098f6bcd4621d373cade4e832627b4f6"
	assert.Equal(t, expectedHash, GetMD5Hash(data))
}

func TestGetSha256Hash(t *testing.T) {
	part := "test"
	key := "mySecretKey"
	expectedHash := "b8f312b17880338cc7855f5d30f6ccec42da0828bce4dc0eeab8ee0777c60af2"
	hash, err := GetSha256Hash(part, key)
	assert.NoError(t, err)
	assert.Equal(t, expectedHash, hash)
}

func TestGetSha256HashNegative(t *testing.T) {
	part := "test"
	key := ""
	expectedError := "empty key"

	_, err := GetSha256Hash(part, key)
	assert.Error(t, err)
	assert.EqualError(t, err, expectedError)
}
