package model

/*
用户模型
*/

import (
	"NightGoBlog/utils/errmsg"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null " json:"username"`
	Password string `gorm:"type:varchar(20);not null " json:"password"`
	Role     int    `gorm:"type:int " json:"role"`
}

// CheckUser 查询用户是否存在
func CheckUser(username string) (code int) {
	var user User
	db.Select("id").Where("username = ?", username).First(&user)
	if user.ID > 0 {
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCSE

}

// CreateUser 新增用户
func CreateUser(user *User) int {
	err := db.Create(&user).Error
	if err != nil {
		return errmsg.ERROR // 500
	}
	return errmsg.SUCCSE

}

// GetUsers GetUser 查询用户列表
func GetUsers(pageSize int, pageNum int) []User {
	var users []User
	err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
	if err != nil {
		return nil
	}
	return users
}
