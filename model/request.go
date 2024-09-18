package model

type Request struct {
	URL     string
	Method  string
	Headers map[string]string
	Body    []byte
	// 其他相關字段
}
