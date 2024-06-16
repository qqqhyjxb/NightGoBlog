package model

/*
用户模型
*/

import (
	"NightGoBlog/utils/errmsg"
	"encoding/base64"
	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
	"log"
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
	// 先对密码进行加密，再添加用户
	//user.Password = ScryptPw(user.Password)
	//上述功能改用BeforeSave钩子函数实现自动调用

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

// EditUser 编辑用户  , 在编辑用户功能中一般不涉及密码的修改，密码修改设计到身份验证应该单独开发一个方法
func EditUser(id int, user *User) int {
	var u User
	var maps = make(map[string]interface{})
	maps["username"] = user.Username
	maps["role"] = user.Role
	err := db.Model(&u).Where("id = ?", id).Updates(maps)
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// DeleteUser 删除用户
func DeleteUser(id int) int {
	var user User
	// 软删除，数据库中并不会真的删除该数据，而是将该记录的DeleteAt设置为当前时间，而后的一般查询方法将无法查找到此条记录
	err := db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// BeforeCreate  ScryptPw 密码加密
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Password = ScryptPw(u.Password)
	return
}

func ScryptPw(password string) string {
	const KeyLen = 10
	salt := make([]byte, 8)
	salt = []byte{12, 32, 4, 6, 66, 22, 222, 11}

	HashPw, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, KeyLen)
	if err != nil {
		log.Fatal(err)
	}

	Fpw := base64.StdEncoding.EncodeToString(HashPw)
	return Fpw
}
