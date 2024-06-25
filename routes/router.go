package routes

import (
	v1 "NightGoBlog/api/v1"
	"NightGoBlog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()
	//r := gin.New()  该方法比上面的Default少了log和错误恢复功能，建议使用Default
	router := r.Group("api/v1")
	{
		// 用户模块的路由接口
		router.POST("user/add", v1.AddUser)
		router.GET("users", v1.GetUsers)
		router.PUT("user/:id", v1.EditUser)
		router.DELETE("user/:id", v1.DeleteUser)
		// 分类模块的路由接口
		router.POST("category/add", v1.AddCategory)
		router.GET("categorys", v1.GetCategorys)
		router.PUT("category/:id", v1.EditCategory)
		router.DELETE("category/:id", v1.DeleteCategory)
		// 文章模块的路由接口
		router.POST("article/add", v1.AddArticle)
		router.GET("articles", v1.GetArt)
		router.GET("article/info/:id", v1.GetArtInfo)
		router.GET("article/list/:id", v1.GetCateArt)
		router.PUT("article/:id", v1.EditArt)
		router.DELETE("article/:id", v1.DeleteArt)

	}

	err := r.Run(utils.HttpPort)
	if err != nil {
		return
	}
}
