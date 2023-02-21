package serializer

// 基础序列化
type Response struct {
	StatusCode int         `json:"status_code"`
	Data       interface{} `json:"data"`
	Msg        string      `json:"msg"`
	Error      error       `json:"error"`
}

// 带 token 的 data
type TokenData struct {
	User  interface{} `json:"user"`
	Token string      `json:"token"`
}

type DataList struct {
	Item  interface{} `json:"item"`
	Total uint        `json:"total"`
}

// list 响应
func BuildListResponse(item interface{}, total int) Response {
	return Response{
		Data: DataList{
			Item:  item,
			Total: uint(total),
		},
	}
}
