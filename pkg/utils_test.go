package goparsy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertToTypeString(t *testing.T) {
	var tests = []struct {
		given string
		want  string
	}{
		{"test", "test"},
		{"Hello World", "Hello World"},
	}
	for _, tt := range tests {
		assert.Equal(t, ConvertToType(tt.given, "STRING"), tt.want)
	}
}

func TestConvertToTypeInteger(t *testing.T) {
	var tests = []struct {
		given string
		want  int64
	}{
		{"12", 12},
		{"123123", 123123},
	}
	for _, tt := range tests {
		assert.Equal(t, ConvertToType(tt.given, "INTEGER"), tt.want)
	}
}

func TestConvertToTypeFloat(t *testing.T) {
	var tests = []struct {
		given string
		want  float64
	}{
		{"12.0", 12.0},
		{"123.123", 123.12300109863281},
	}
	for _, tt := range tests {
		assert.Equal(t, ConvertToType(tt.given, "FLOAT"), tt.want)
	}
}
