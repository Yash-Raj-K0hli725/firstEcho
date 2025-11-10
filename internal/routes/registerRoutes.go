package routes

import (
	"FirstGo/internal/handlers"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	e.GET("/", handlers.Welcome)
}
