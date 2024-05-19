package serve

import "github.com/gin-gonic/gin"

type Router interface {
	BindOn(r *gin.Engine)
}
