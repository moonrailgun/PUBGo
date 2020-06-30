package config

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestDefaultConfig(t *testing.T) {
	assert.Equal(t, "pubgo", Config.DB.Name)
	assert.Equal(t, "localhost", Config.DB.Host)
	assert.Equal(t, "3306", Config.DB.Port)
}

func TestGetCurrentPath(t *testing.T) {
	wd, _ := os.Getwd()
	assert.Equal(t, getCurrentPath(), wd)
}
