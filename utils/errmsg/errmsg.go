package errmsg

/*
错误信息
*/

// 声明状态码常量
const (
	SUCCSE = 200
	ERROR  = 500

	// code = 1000... 用户模块的错误
	ERROR_USERNAME_USED    = 1001 //用户已存在
	ERROR_PASSWRD_WRONG    = 1002 //密码错误
	ERROR_USER_NOT_EXIST   = 1003 //用户不存在
	ERROR_TOKEN_EXIST      = 1004 //用户携带token不存在
	ERROR_TOKEN_RUNTIME    = 1005 //用户token过期
	ERROR_TOKEN_WRONG      = 1006 //用户token错误
	ERROR_TOKEN_TYPE_WRONG = 1007 //用户token格式错误
	// code = 2000... 文章模块的错误

	// code = 3000... 分类模块的错误
	ERROR_CATENAME_USED = 2001 //分类已存在
)

// CodeMsg 声明字典
var CodeMsg = map[int]string{
	SUCCSE:                 "OK",
	ERROR:                  "FAIL",
	ERROR_USERNAME_USED:    "用户名已存在！",
	ERROR_PASSWRD_WRONG:    "密码错误！",
	ERROR_USER_NOT_EXIST:   "用户不存在！",
	ERROR_TOKEN_EXIST:      "TOKEN不存在",
	ERROR_TOKEN_RUNTIME:    "TOKEN已过期",
	ERROR_TOKEN_WRONG:      "TOKEN不正确",
	ERROR_TOKEN_TYPE_WRONG: "TOKEN格式错误",
	ERROR_CATENAME_USED:    "分类已存在",
}

// GetErrMsg 输出错误信息的方法
func GetErrMsg(code int) string {
	return CodeMsg[code]
}
