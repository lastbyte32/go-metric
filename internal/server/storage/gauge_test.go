package storage

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//func TestGaugeMemStorage_All(t *testing.T) {
//	type fields struct {
//		values map[string]float64
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		want   map[string]float64
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			g := &GaugeMemStorage{
//				values: tt.fields.values,
//			}
//			assert.Equalf(t, tt.want, g.All(), "All()")
//		})
//	}
//}
//
//func TestGaugeMemStorage_Get(t *testing.T) {
//	type fields struct {
//		values map[string]float64
//	}
//	type args struct {
//		name string
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//		want   float64
//		want1  bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			g := &GaugeMemStorage{
//				values: tt.fields.values,
//			}
//			got, got1 := g.Get(tt.args.name)
//			assert.Equalf(t, tt.want, got, "Get(%v)", tt.args.name)
//			assert.Equalf(t, tt.want1, got1, "Get(%v)", tt.args.name)
//		})
//	}
//}

func TestGaugeMemStorage_Update(t *testing.T) {

	type fields struct {
		values map[string]float64
	}

	type args struct {
		name  string
		value float64
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
	}{
		{
			"Test1 update and get",
			fields{map[string]float64{"test": 99.99}},
			args{"test", 3},
			3,
		},

		{
			"Test2 override value by key",
			fields{map[string]float64{"test": 100.99}},
			args{"test2", 99.99},
			99.99,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GaugeMemStorage{
				values: tt.fields.values,
			}
			g.Update(tt.args.name, tt.args.value)

			value, exist := g.Get(tt.args.name)
			assert.True(t, exist)
			assert.Equal(t, tt.want, value)

		})
	}
}

func TestNewGaugeMemStorage(t *testing.T) {
	tests := []struct {
		name string
		want GaugeStorage
	}{
		{"First",
			NewGaugeMemStorage(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NotNil(t, tt.want, "NewGaugeMemStorage() returned nil")
			assert.Equalf(t, tt.want, NewGaugeMemStorage(), "NewGaugeMemStorage()")
		})
	}
}
