package controllers

import (
	"net/http"

	"github.com/RohitKMishra/google_oauth2_golang/models"
	"github.com/gin-gonic/gin"
)

func GetMe(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"user": models.FilteredResponse(&currentUser)}})
}
