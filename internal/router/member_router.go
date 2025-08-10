package router

import (
	"github.com/fredianto2405/catetin-api/internal/member"
	"github.com/fredianto2405/catetin-api/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterMemberRoutes(rg *gin.RouterGroup, handler *member.Handler) {
	rg.Use(middleware.JWTAuthMiddleware())
	rg.POST("", handler.CreateMember)
	rg.GET("", handler.GetMembers)
	rg.PUT("/:id", handler.UpdateMember)
	rg.DELETE("/:id", handler.DeleteMember)
}
