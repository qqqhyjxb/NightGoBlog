package v1

/*
user模块的接口
*/

import (
	"NightGoBlog/model"
	"NightGoBlog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// AddUser 添加用户
func AddUser(c *gin.Context) {
	var user model.User
	// 绑定JSON模型
	_ = c.ShouldBindJSON(&user)
	// 检查用户是否存在
	code := model.CheckUser(user.Username)
	// 不存在则添加用户
	if code == errmsg.SUCCSE {
		model.CreateUser(&user)
	} else if code == errmsg.ERROR_USERNAME_USED {
		code = errmsg.ERROR_USERNAME_USED
	}

	// 响应
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    user,
		"message": errmsg.GetErrMsg(code),
	})

}

// 查询单个用户

// GetUsers 查询用户列表
func GetUsers(c *gin.Context) {
	// 接收的同时要将字符串转换为整型
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	// 传入-1 则取消分页功能
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}

	data := model.GetUsers(pageSize, pageNum)
	code := errmsg.SUCCSE
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// EditUser 编辑用户
func EditUser(c *gin.Context) {
	// todo 编辑
}

// DeleteUser 删除用户
func DeleteUser(c *gin.Context) {
	// todo 删除用户
}
