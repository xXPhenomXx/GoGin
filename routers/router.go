package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	_ "GoGin/docs"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"GoGin/middleware/jwt"
	"GoGin/middleware/cors"
	"GoGin/pkg/export"
	"GoGin/pkg/setting"
	"GoGin/pkg/upload"
	"GoGin/routers/api"
	"GoGin/routers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(cors.CORS())

	r.Use(gin.Recovery())
	gin.SetMode(setting.ServerSetting.RunMode)

	r.StaticFS("/export", http.Dir(export.GetExcelFullPath()))
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))

	r.GET("/auth", api.CheckAuth)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/upload", api.UploadImage)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		// user routes
		apiv1.GET("/users", v1.GetUsers)


	}

	return r
}
