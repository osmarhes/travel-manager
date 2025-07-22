package server

import (
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/osmarhes/travel-manager/internal/auth"
	"github.com/osmarhes/travel-manager/internal/authmiddleware"
	"github.com/osmarhes/travel-manager/internal/travel"
	"github.com/osmarhes/travel-manager/internal/user"
	"github.com/osmarhes/travel-manager/pkg/database"
)

func Run() {
	database.Connect()
	db := database.DB

	userRepo := user.NewRepository(db)
	authHandler := auth.NewHandler(userRepo)

	travelRepo := travel.NewRepository(db)
	travelHandler := travel.NewHandler(travelRepo)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/auth/register", authHandler.Register)
	e.POST("/auth/login", authHandler.Login)

	t := e.Group("/travels")
	t.Use(authmiddleware.JWTMiddleware)
	t.POST("", travelHandler.Create)
	t.GET("/:id", travelHandler.GetByID)
	t.GET("", travelHandler.List)
	t.PUT("/:id/status", travelHandler.UpdateStatus)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server running on port", port)
	e.Logger.Fatal(e.Start(":" + port))
}
