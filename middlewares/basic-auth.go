package middlewares

import (
	"os"

	"github.com/gin-gonic/gin"
)

func BasicAuth() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		os.Getenv("BASIC_AUTH_LOGIN"): os.Getenv("BASIC_AUTH_PASSWORD"),
	})
}
