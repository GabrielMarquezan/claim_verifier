package router

import (
	"api/config"
	"api/controller"
	"api/repository"
	"context"

	"github.com/jmoiron/sqlx"
)

type Injection struct {
	db  *sqlx.DB
	cfg *config.Config
	ctx *context.Context
}

func (injection *Injection) NewClaimController() *controller.ClaimController {
	r := &repository.ClaimRepository{DB: injection.db, Ctx: injection.ctx}
	return &controller.ClaimController{ClaimRepository: r, Cfg: injection.cfg}
}
