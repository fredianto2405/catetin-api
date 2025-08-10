package router

import (
	"github.com/fredianto2405/catetin-api/internal/auth"
	"github.com/fredianto2405/catetin-api/pkg/errors"
	"github.com/fredianto2405/catetin-api/pkg/logger"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"time"
)

func SetupRouter(db *sqlx.DB) *gin.Engine {
	// init validator
	errors.InitValidator()

	// init logger
	logger.InitLogger()

	r := gin.Default()

	// middleware cors
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// error handler
	r.Use(errors.ErrorHandler())

	// auth routes
	authRepo := auth.NewRepository(db)
	authService := auth.NewService(authRepo)
	authHandler := auth.NewHandler(authService)
	authGroup := r.Group("/api/v1/auth")
	RegisterAuthRoutes(authGroup, authHandler)

	return r
}
