package utils

import (
	"github.com/json-iterator/go/extra"
	"unicode"
)

func init() {
	// 设置jsoniter的命名规范
	extra.SetNamingStrategy(LowerCaseFirst)
}

// 首字母小写
func LowerCaseFirst(name string) string {
	var newName []rune

	for i, c := range name {
		if i == 0 {
			newName = append(newName, unicode.ToLower(c))
		} else {
			newName = append(newName, c)
		}
	}
	return string(newName)
}
