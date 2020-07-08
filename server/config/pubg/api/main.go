package api

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"github.com/moonrailgun/PUBGo/server/config/db"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
)

type API struct {
	Key string
	Url string
}

type ShardType string

const (
	STEAM  ShardType = "steam"
	KAKAO  ShardType = "kakao"
	PSN    ShardType = "psn"
	STADIA ShardType = "stadia"
	XBOX   ShardType = "xbox"
)

func NewAPI(key string) *API {
	return &API{
		Key: key,
		Url: "https://api.playbattlegrounds.com",
	}
}

// http 请求
func httpRequest(url, key string) (*bytes.Buffer, error) {
	logrus.WithField("url", url).Info("pubg api request")

	// Create request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Set request options
	req.Header.Set("Authorization", key)
	req.Header.Set("Accept", "application/vnd.api+json")
	req.Header.Set("Accept-Encoding", "gzip")

	// Execute request
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	// Check http response code
	if response.StatusCode != 200 {
		err := response.Body.Close()
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("HTTP request failed: %s", response.Status)
	}

	// Retrieve response body
	var reader io.ReadCloser
	switch response.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(response.Body)
		if err != nil {
			return nil, err
		}
	default:
		reader = response.Body
	}
	defer reader.Close()

	var buffer bytes.Buffer
	buffer.ReadFrom(reader)

	// 增加记录
	// 注意此处为了不引入数据库模型(会循环引用) 使用了raw sql
	db.GetDb().Exec(`INSERT INTO model_pubg_request_log(url, resp, created_at) VALUES (?, ?, now())`, url, buffer.String())

	return &buffer, nil
}
