package e

var MsgFlags = map[int]string{
	Success:       "ok",
	Error:         "fail",
	InvaildParams: "Parameter error.",

	ErrorExistUser:         "User has been registered.",
	ErrorFailEncryption:    "Password encryption failed.",
	ErrorExistUserNotFound: "User does not exist.",
	ErrorNotCompare:        "Password error.",
	ErrorAuthToken:         "Token authentication failed.",

	StatusNotFound:             "Status not found.",
	ErrorAuthCheckTokenFail:    "Token validation failed",
	ErrorAuthCheckTokenTimeout: "Token has expired.",
}

// 获取状态码对应的信息
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if !ok {
		return MsgFlags[Error]
	}
	return msg
}
