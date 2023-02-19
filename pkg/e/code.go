package e

const (
	Success       = 0
	Error         = -1
	InvaildParams = 400

	ErrorExistUser         = 30001
	ErrorFailEncryption    = 30002
	ErrorExistUserNotFound = 30003
	ErrorNotCompare        = 30004
	ErrorAuthToken         = 30005

	StatusNotFound             = 404
	ErrorAuthCheckTokenFail    = 405
	ErrorAuthCheckTokenTimeout = 406
)
