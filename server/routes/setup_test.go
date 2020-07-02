package routes

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// 发送测试请求
func testRequest(method string, url string, body io.Reader) *httptest.ResponseRecorder {
	router := SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(method, url, body)
	router.ServeHTTP(w, req)

	return w
}

func TestPingRoute(t *testing.T) {
	w := testRequest("GET","/ping", nil)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

func TestGetPlayerInfo(t *testing.T) {
	w := testRequest("GET", "/player/shard/steam/name/WackyJacky101", nil)

	assert.Equal(t, 200, w.Code)
	fmt.Printf(w.Body.String())
}
