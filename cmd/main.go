package main

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"Mini_Project/config"
	"Mini_Project/internal/app/handler"
	"Mini_Project/internal/app/middleware"
	"Mini_Project/internal/app/repository"
	"Mini_Project/internal/app/service"
	"Mini_Project/pkg/logger"
)

func main() {
	// Load config
	config.Init()

	// Setup logger
	logger.Init()

	// Setup database connection
	db, err := gorm.Open(mysql.Open(viper.GetString("db.dsn")), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	if err := db.AutoMigrate(&model.User{}, &model.Meteran{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	// Setup Echo instance
	e := echo.New()

	// Setup middleware
	authMiddleware := middleware.NewAuthMiddleware(config.GetJWTSecret())

	// Setup repositories
	userRepo := repository.NewUserRepository(db)
	meteranRepo := repository.NewMeteranRepository(db)

	// Setup services
	userService := service.NewUserService(userRepo)
	meteranService := service.NewMeteranService(meteranRepo)

	// Setup handlers
	authHandler := handler.NewAuthHandler(config.GetJWTSecret())
	userHandler := handler.NewUserHandler(userService)
	meteranHandler := handler.NewMeteranHandler(meteranService)

	// Setup routes
	authGroup := e.Group("/auth")
	authGroup.POST("/login", authHandler.Login)

	userGroup := e.Group("/users")
	userGroup.Use(authMiddleware)
	userGroup.GET("", userHandler.GetUsers)
	userGroup.GET("/:id", userHandler.GetUser)
	userGroup.POST("", userHandler.CreateUser)
	userGroup.PUT("/:id", userHandler.UpdateUser)
	userGroup.DELETE("/:id", userHandler.DeleteUser)

	meteranGroup := e.Group("/meterans")
	meteranGroup.Use(authMiddleware)
	meteranGroup.GET("", meteranHandler.GetMeterans)
	meteranGroup.GET("/:id", meteranHandler.GetMeteran)
	meteranGroup.POST("", meteranHandler.CreateMeteran)
	meteranGroup.PUT("/:id", meteranHandler.UpdateMeteran)
	meteranGroup.DELETE("/:id", meteranHandler.DeleteMeteran)

	// Start server
	addr := fmt.Sprintf(":%v", viper.GetString("app.port"))
	if err := e.Start(addr); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
