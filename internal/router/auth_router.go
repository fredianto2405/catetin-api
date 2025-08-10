package router

import (
	"github.com/fredianto2405/catetin-api/internal/auth"
	"github.com/fredianto2405/catetin-api/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(rg *gin.RouterGroup, handler *auth.Handler) {
	rg.POST("/login", handler.Login)
	rg.POST("/change-password", middleware.JWTAuthMiddleware(), handler.ChangePassword)
}
