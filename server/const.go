package server

// ERROR CODE

const (
	ERR_CODE_PARAMS = 1
	ERR_CODE_SYSTEM
	ERR_CODE_NOT_FOUND
)

// ERROR MESSAGE
const (
	ERR_MSG_SYSTEM        = "系统错误"
	ERR_MSG_PARAMS        = "参数错误"
	ERR_MSG_NOT_FOUND     = "未找到该记录"
	ERR_MSG_GET_SESSION   = "获取session key错误"
	ERR_MSG_GET_USER_INFO = "获取user info错误"
)
