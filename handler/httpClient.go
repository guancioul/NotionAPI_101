package handler

import (
	"bytes"
	"net/http"

	"github.com/swaggo/swag/example/celler/model"
)

type Client struct {
	httpClient *http.Client
}

func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{},
	}
}

// SendRequest 通用的 HTTP 請求方法
func (c *Client) SendRequest(req model.Request) (*http.Response, error) {
	// 在此處實現請求的發送邏輯，並返回響應
	// 可以在這裡處理例如設置請求頭、處理請求錯誤等等
	// 這個方法可以支持不同的 HTTP 方法

	httpRequest, err := http.NewRequest(req.Method, req.URL, bytes.NewBuffer(req.Body))
	if err != nil {
		return nil, err
	}

	// 設置請求頭
	for key, value := range req.Headers {
		httpRequest.Header.Set(key, value)
	}

	// 發送請求
	response, err := c.httpClient.Do(httpRequest)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Get 方法用於發送 GET 請求
func (c *Client) Get(url string, headers map[string]string) (*http.Response, error) {
	req := model.Request{
		URL:     url,
		Method:  "GET",
		Headers: headers,
	}
	return c.SendRequest(req)
}

// Post 方法用於發送 POST 請求
func (c *Client) Post(url string, headers map[string]string, body []byte) (*http.Response, error) {
	req := model.Request{
		URL:     url,
		Method:  "POST",
		Headers: headers,
		Body:    body,
	}
	return c.SendRequest(req)
}

// Put 方法用於發送 PUT 請求
func (c *Client) Put(url string, headers map[string]string, body []byte) (*http.Response, error) {
	req := model.Request{
		URL:     url,
		Method:  "PUT",
		Headers: headers,
		Body:    body,
	}
	return c.SendRequest(req)
}

// Patch 方法用於發送 PATCH 請求
func (c *Client) Patch(url string, headers map[string]string, body []byte) (*http.Response, error) {
	req := model.Request{
		URL:     url,
		Method:  "PATCH",
		Headers: headers,
		Body:    body,
	}
	return c.SendRequest(req)
}
