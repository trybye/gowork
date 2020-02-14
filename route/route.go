package route

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"demoproject/middleware"
	"demoproject/route/api"
)

func NewRoute() *gin.Engine {
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//r.Use(middleware.Cors())
	r.Use(cors.Default()) //跨域
	v1 := r.Group("/") //no cookie
	{
		v1.GET("demo", api.DemoApi)

		authv1 := v1.Group("/")
		authv1.Use(middleware.AuthLogin)

		{
			//for test
			authv1.POST("auth_demo", api.DemoApi) //c客户申请充值
		}

	}

	return r
}
