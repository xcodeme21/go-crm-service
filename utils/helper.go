package helper

import (
	"github.com/gin-gonic/gin"
)

func JSONResponse(ctx *gin.Context, status int, data interface{}, errorMessage string) {
	ctx.JSON(status, gin.H{
		"status":        status,
		"data":          data,
		"error_message": errorMessage,
	})
}
