package controller

import (
	"api/model"
	"fmt"

	"github.com/gin-gonic/gin"
)

func CreateClaim(context *gin.Context) {
	claim := model.Claim{}
	context.ShouldBindJSON(&claim)

	fmt.Println("JSON:", claim)
}
