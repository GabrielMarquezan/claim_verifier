package controller

import (
	"api/model"
	"api/repository"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

type ClaimController struct {
	ClaimRepository *repository.ClaimRepository
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
}
