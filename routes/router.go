package routes

import (
	v1 "NightGoBlog/api/v1"
	"NightGoBlog/middleware"
	"NightGoBlog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()
	//r := gin.New()  该方法比上面的Default少了log和错误恢复功能，建议使用Default
	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())
	{
		// 用户模块的路由接口
		auth.PUT("user/:id", v1.EditUser)
		auth.DELETE("user/:id", v1.DeleteUser)
		// 分类模块的路由接口
		auth.POST("category/add", v1.AddCategory)
		auth.PUT("category/:id", v1.EditCategory)
		auth.DELETE("category/:id", v1.DeleteCategory)
		// 文章模块的路由接口
		auth.POST("article/add", v1.AddArticle)
		auth.PUT("article/:id", v1.EditArt)
		auth.DELETE("article/:id", v1.DeleteArt)
	}
	// 不需要权限的
	router := r.Group("api/v1")
	{
		router.POST("user/add", v1.AddUser)
		router.GET("users", v1.GetUsers)
		router.GET("categorys", v1.GetCategorys)
		router.GET("articles", v1.GetArt)
		router.GET("article/info/:id", v1.GetArtInfo)
		router.GET("article/list/:id", v1.GetCateArt)
		router.POST("login", v1.Login)
	}

	err := r.Run(utils.HttpPort)
	if err != nil {
		return
	}
}
