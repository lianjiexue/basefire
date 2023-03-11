package bootstrap

import (
	"app/config"
	"app/internal/repo"
	"app/routes"
	"github.com/gin-gonic/gin"
)

type Application struct {
	Name string
	Port string
}

func (app *Application) Start() {
	//0.初始化配置
	config.NewConfig()
	//1.定义db
	repo.NewDb()
	//2.加载路由
	router := gin.Default()
	routes.NewApiRouter(router)
	routes.NewWebRouter(router)
	//3.启动gin
	router.Run(app.Port) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
