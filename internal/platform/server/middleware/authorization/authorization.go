package authorization

import (
	"github.com/alexperezortuno/go-auth/common"
	"github.com/alexperezortuno/go-auth/internal/platform/storage/auth"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var header = c.GetHeader("Authorization")

		if header == "" {
			log.Printf("[Middleware] %s Unauthorized access attempt", c.Request.RemoteAddr)
			c.AbortWithStatusJSON(http.StatusUnauthorized, common.ErrorResponse("credentials is not provided"))
			return
		}

		_, err := auth.ValidateToken(header)
		if err != "" {
			log.Printf("[Middleware] %s Unauthorized access attempt", c.Request.RemoteAddr)
			c.AbortWithStatusJSON(http.StatusUnauthorized, common.ErrorResponse(err))
			return
		}

		c.Next()
	}
}
