package util

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

type HttpClient struct {
	baseURL    string
	httpClient *http.Client
	headers    map[string]string
}

func NewHttpClient(timeout time.Duration) *HttpClient {
	return &HttpClient{
		httpClient: &http.Client{
			Timeout: timeout * time.Second, // 延长超时时间
			Transport: &http.Transport{
				MaxIdleConns:       100,
				IdleConnTimeout:    90 * time.Second,
				DisableCompression: false,
				DisableKeepAlives:  false,
			},
		},
		headers: make(map[string]string),
	}
}

func (c *HttpClient) SetHeader(key, value string) {
	c.headers[key] = value
}

func (c *HttpClient) PostJSON(ctx context.Context, path string, body interface{}) ([]byte, error) {
	fullURL := c.baseURL + path
	log.Printf("请求地址---------->%s", fullURL)
	log.Printf("请求参数---------->%s", body)
	jsonData, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", fullURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	for k, v := range c.headers {
		req.Header.Set(k, v)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
