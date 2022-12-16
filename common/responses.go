package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorResponse(err string) gin.H {
	return gin.H{"message": err, "status": false}
}

func SuccessResponse(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{"result": data, "status": true, "message": "success"})
}

func SuccessResponseWithMessage(ctx *gin.Context, data interface{}, message string) {
	ctx.JSON(http.StatusOK, gin.H{"result": data, "status": true, "message": message})
}

func BadRequest(ctx *gin.Context, err string) {
	ctx.JSON(http.StatusBadRequest, gin.H{"message": err, "status": false})
}

func Unauthorized(ctx *gin.Context, err string) {
	ctx.JSON(http.StatusUnauthorized, gin.H{"message": err, "status": false})
}

func Forbidden(ctx *gin.Context, err string) {
	ctx.JSON(http.StatusForbidden, gin.H{"message": err, "status": false})
}

func NotFound(ctx *gin.Context, err string) {
	ctx.JSON(http.StatusNotFound, gin.H{"message": err, "status": false})
}

func InternalServerError(ctx *gin.Context, err string) {
	ctx.JSON(http.StatusInternalServerError, gin.H{"message": err, "status": false})
}
