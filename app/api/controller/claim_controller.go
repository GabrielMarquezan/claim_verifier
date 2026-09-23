package controller

import (
	"api/config"
	"api/model"
	"api/repository"
	"api/service"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

type ClaimController struct {
	ClaimRepository *repository.ClaimRepository
	Cfg             *config.Config
}

func (controller *ClaimController) CreateClaim(context *gin.Context) {
	claim := model.Claim{}
	context.ShouldBindJSON(&claim)

	controller.ClaimRepository.Create(&claim)

	jsonClaim, err := json.Marshal(claim)
	if err != nil {
		context.Abort()
	}

	context.Writer.Write(jsonClaim)

	err = service.StartClaimAnalysis(jsonClaim, controller.Cfg.RAGWorkerAddr)
	if err != nil {
		context.Abort()
	}
}
