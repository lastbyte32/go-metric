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
