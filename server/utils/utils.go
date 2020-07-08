package utils

import (
	jsoniter "github.com/json-iterator/go"
)

func QuickMarshal(jsonData interface{}) string {
	str, _ := jsoniter.Marshal(jsonData)
	return string(str[:])
}
