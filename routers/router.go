package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "github.com/xxphenomxx/GoGin/docs"

	"github.com/xxphenomxx/GoGin/middleware/cors"
	"github.com/xxphenomxx/GoGin/middleware/jwt"
	"github.com/xxphenomxx/GoGin/pkg/export"
	"github.com/xxphenomxx/GoGin/pkg/setting"
	"github.com/xxphenomxx/GoGin/pkg/upload"
	"github.com/xxphenomxx/GoGin/routers/api"
	//"github.com/xxphenomxx/GoGin/routers/api/v1"
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

	// protected routes
	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		// user routes
		apiv1.GET("/users/all", v1.GetUsers)
	}

	return r
}
