package utils

import "encoding/json"

func QuickMarshal(jsonData interface{}) string {
	str, _ := json.Marshal(jsonData)
	return string(str[:])
}
