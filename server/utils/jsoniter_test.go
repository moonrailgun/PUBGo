package utils

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJsoniterNamingStrategy(t *testing.T) {
	str1, _ := jsoniter.MarshalToString(map[string]int{
		"a": 1,
		"b": 2,
	})

	assert.Equal(t, `{"a":1,"b":2}`, str1)

	str2, _ := jsoniter.MarshalToString(struct {
		UserName      string
		FirstLanguage string
	}{
		UserName:      "taowen",
		FirstLanguage: "Chinese",
	})
	assert.Equal(t, `{"userName":"taowen","firstLanguage":"Chinese"}`, str2)
}
