package services

import (
	"net/http"
	"vox/vox/initializers"

	"github.com/gin-gonic/gin"
)

func GetAll(ctx *gin.Context) {

	db := initializers.DB
	result := map[string]interface{}{}
	db.Table("users").Take(&result)

	ctx.JSON(http.StatusOK, gin.H{
		"deneme": result,
	})

}
