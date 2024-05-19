package serve

import (
	_ "q6/docs" // load swagger docs

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var _ Router = &SwaggerRouter{}

type SwaggerRouter struct {
}

func NewSwaggerRouter() Router {
	return &SwaggerRouter{}
}

func (r *SwaggerRouter) BindOn(e *gin.Engine) {
	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
