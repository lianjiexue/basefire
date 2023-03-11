package routes

import (
	"app/internal/http/api/index"
	"app/internal/http/api/login"
	"app/internal/http/api/news"
	"app/internal/middleware"
	"github.com/gin-gonic/gin"
)

func NewApiRouter(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.GET("/index", index.Index)

		api.POST("/login", login.Login)
		api.POST("/wx/login", login.WxLogin)

		api.GET("/news", news.Lists)
		api.GET("/news/detail/{id}", news.Detail)
		api.GET("/bind", index.Index)
		api.GET("/product/info/{id}", index.Index)
		api.GET("/product/lists", index.Index)
		api.GET("/category/lists", index.Index)
		api.GET("/category/product/lists", index.Index)

		api.Use(middleware.VerifyToken())
		{
			api.GET("/cart/lists", index.Index)
			api.GET("/cart/info", index.Index)
			api.GET("/cart/add", index.Index)
			api.GET("/cart/del", index.Index)

			api.GET("/order/submit", index.Index)
			api.GET("/order/product/set", index.Index)
			api.GET("/order/product/get", index.Index)
			api.GET("/order/lists", index.Index)
			api.GET("/order/info", index.Index)
			api.GET("/order/create", index.Index)
			api.GET("/order/edit", index.Index)

			api.GET("/user/info", index.Index)
			api.GET("/user/address/add", index.Index)
			api.GET("/user/address/edit", index.Index)
			api.GET("/user/address/del", index.Index)
			api.GET("/user/address/default", index.Index)
			api.GET("/user/address/list", index.Index)
		}
	}
}
