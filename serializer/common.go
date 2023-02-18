package serializer

// 基础序列化
type Response struct {
	StatusCode int         `json:"status_code"`
	Data       interface{} `json:"data"`
	Msg        string      `json:"msg"`
	Error      error       `json:"error"`
}
