package router

import (
	"api/controller"
	"api/repository"
	"context"

	"github.com/jmoiron/sqlx"
)

type Injection struct {
	db  *sqlx.DB
	ctx *context.Context
}

func (injection *Injection) NewClaimController() *controller.ClaimController {
	r := &repository.ClaimRepository{DB: injection.db, Ctx: injection.ctx}
	return &controller.ClaimController{ClaimRepository: r}
}
