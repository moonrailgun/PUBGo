package models

import (
	"github.com/moonrailgun/PUBGo/server/config/db"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSetGlobalConfig(t *testing.T) {
	err := SetGlobalConfig("Test Name", "Test Value")
	if err != nil {
		panic(err)
	}

	m := new(GlobalConfig)
	err2 := db.GetDb().Where("name = ?", "Test Name").First(m).Error
	if err2 != nil {
		panic(err2)
	}
	assert.Equal(t, "Test Value", m.Value)
}

func TestGetGlobalConfig(t *testing.T) {
	err := SetGlobalConfig("Test Name2", "Test Value2")
	if err != nil {
		panic(err)
	}

	value := GetGlobalConfig("Test Name2")
	assert.Equal(t, "Test Value2", value)
}
