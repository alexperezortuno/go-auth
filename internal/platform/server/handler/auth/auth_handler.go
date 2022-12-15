package auth

import (
	"github.com/alexperezortuno/go-auth/internal/platform/storage/auth"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func LoginHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req auth.AuthRequest

		if err := ctx.BindJSON(&req); err != nil {
			log.Printf("[ERROR] %s", err.Error())
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		response, err := auth.ValidateUser(req)
		if err != "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status":  false,
				"message": err,
			})
			return
		}

		ctx.JSON(http.StatusOK, response)
	}
}

func CreateUserHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req auth.UserRequest

		if err := ctx.BindJSON(&req); err != nil {
			log.Printf("[ERROR] %s", err.Error())
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		response, err := auth.CreateUser(req)
		if err != "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status":  false,
				"message": err,
			})
			return
		}

		ctx.JSON(http.StatusOK, response)
	}
}

func ValidateTokenHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var header = ctx.GetHeader("Authorization")

		if header == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status":  false,
				"message": "credentials is required",
			})
			return
		}

		response, err := auth.ValidateToken(header)
		if err != "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status":  false,
				"message": response.Message,
			})
			return
		}

		ctx.JSON(http.StatusOK, response)
	}
}

func RefreshTokenHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req auth.TokenRequest

		if err := ctx.BindJSON(&req); err != nil {
			log.Printf("[ERROR] %s", err.Error())
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		response, err := auth.RefreshToken(req.Token)
		if err != "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status":  false,
				"message": err,
			})
			return
		}

		ctx.JSON(http.StatusOK, response)
	}
}
