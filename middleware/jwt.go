package middleware

import (
	"NightGoBlog/utils"
	"NightGoBlog/utils/errmsg"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

/*
	jwt认证
*/

var JwtKey = []byte(utils.JwtKey)

type MyClaims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

var code int

// SetToken 生成token
func SetToken(username string, password string) (string, int) {
	expireTime := time.Now().Add(10 * time.Hour)
	setClaims := &MyClaims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "nightGoBlog",
		},
	}
	reqClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, setClaims)
	token, err := reqClaim.SignedString(JwtKey)
	if err != nil {
		return "", errmsg.ERROR
	}
	return token, errmsg.SUCCSE
}

// CheckToken 验证token
func CheckToken(token string) (*MyClaims, int) {
	settoken, _ := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if key, code := settoken.Claims.(*MyClaims); code && settoken.Valid {
		return key, errmsg.SUCCSE
	} else {
		return nil, errmsg.ERROR
	}

}

// JwtToken jwt中间件
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := c.Request.Header.Get("NightGoBlog-Token")
		code = errmsg.SUCCSE
		if tokenHeader == "" {
			code = errmsg.ERROR_TOKEN_EXIST
			//c.Abort()
			//return
		}
		checkToken := strings.SplitN(tokenHeader, " ", 2)
		if len(checkToken) != 2 && checkToken[0] != "Bearer" {
			code = errmsg.ERROR_TOKEN_TYPE_WRONG
		}
		key, Tcode := CheckToken(checkToken[1])
		if Tcode != errmsg.SUCCSE {
			code = errmsg.ERROR_TOKEN_WRONG
			c.Abort()
		}
		if time.Now().Unix() > key.ExpiresAt {
			code = errmsg.ERROR_TOKEN_RUNTIME
			c.Abort()
		}
		c.JSON(http.StatusOK, gin.H{
			"code":    code,
			"message": errmsg.GetErrMsg(code),
		})
		c.Set("username", key.Username)
		c.Next()
	}
}
