package router

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func Init(ctx *context.Context, database *sqlx.DB) *gin.Engine {
	router := gin.Default()

	injection := Injection{db: database, ctx: ctx}
	claimController := injection.NewClaimController()

	router.POST("/claims", claimController.CreateClaim)

	return router
}
