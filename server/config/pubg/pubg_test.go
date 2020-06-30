package pubg

import (
	"github.com/moonrailgun/PUBGo/server/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPubgApiKey(t *testing.T) {
	assert.Equal(t, config.Config.PUBG.API, Api.Key)
}