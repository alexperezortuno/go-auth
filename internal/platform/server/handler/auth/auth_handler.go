package auth

import (
	"github.com/alexperezortuno/go-auth/common"
	"github.com/alexperezortuno/go-auth/internal/platform/storage/auth"
	"github.com/alexperezortuno/go-auth/internal/platform/storage/user"
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

func GetUserHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req auth.UserNameRequest

		if err := ctx.BindJSON(&req); err != nil {
			log.Printf("[ERROR] %s", err.Error())
			common.BadRequest(ctx, "username is required")
			return
		}

		u, err := auth.GetUser(req)
		if err != "" {
			common.Unauthorized(ctx, err)
			return
		}

		var res = user.Response{
			ID:       u.ID,
			IdCard:   u.IdCard,
			FullName: u.FullName,
			Name:     u.Name,
			Nickname: u.Nickname,
			LastName: u.LastName,
			Email:    u.Email,
		}

		common.SuccessResponse(ctx, res)
	}
}

func ValidateTokenHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var header = ctx.GetHeader("Authorization")

		if header == "" {
			common.Unauthorized(ctx, "credentials is required")
			return
		}

		response, err := auth.ValidateToken(header)
		if err != "" {
			common.Unauthorized(ctx, response.Message)
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
			common.BadRequest(ctx, err.Error())
			return
		}

		response, err := auth.RefreshToken(req.Token)
		if err != "" {
			common.Unauthorized(ctx, err)
			return
		}

		common.SuccessResponse(ctx, response)
	}
}
