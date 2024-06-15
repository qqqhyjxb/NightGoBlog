package routes

import (
	"NightGoBlog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()
	//r := gin.New()  该方法比上面的Default少了log和错误恢复功能，建议使用Default
	route := r.Group("api/v1")
	{
		route.GET("hello", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"message": "hello world",
			})
		})
	}

	err := r.Run(utils.HttpPort)
	if err != nil {
		return
	}
}
