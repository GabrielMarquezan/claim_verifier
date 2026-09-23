package router

import (
	"api/config"
	"context"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func Init(ctx *context.Context, database *sqlx.DB, cfg *config.Config) *gin.Engine {
	router := gin.Default()

	injection := Injection{db: database, ctx: ctx, cfg: cfg}
	claimController := injection.NewClaimController()

	router.POST("/claims", claimController.CreateClaim)

	return router
}
