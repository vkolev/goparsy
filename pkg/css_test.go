package goparsy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvert(t *testing.T) {
	var tests = []struct {
		given string
		want  string
	}{
		{"#sample", `/descendant-or-self::*/*[@id="sample"]`},
		{"div.test", `/descendant-or-self::*/*[self::div and contains(concat(" ", @class, " "), " test ")]`},
	}
	for _, tt := range tests {
		assert.Equal(t, Convert(tt.given, GLOBAL), tt.want)
	}
}
