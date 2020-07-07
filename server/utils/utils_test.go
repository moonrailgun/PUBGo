package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQuickMarshal(t *testing.T) {
	ret := QuickMarshal(map[string]int{
		"a": 1,
		"b": 2,
	})

	assert.Equal(t, `{"a":1,"b":2}`, ret)
}
