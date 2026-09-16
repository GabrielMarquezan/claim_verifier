package router

import (
	"api/controller"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	router := gin.Default()

	router.POST("/claims", controller.CreateClaim)

	return router
}
