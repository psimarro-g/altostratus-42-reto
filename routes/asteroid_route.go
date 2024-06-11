package routes

import (
	"altostratus-42-reto/controllers"

	"github.com/labstack/echo/v4"
)

func AsteroidRoute(e *echo.Echo) {
	api := e.Group("/api/v1", serverHeader)

	// Asteroid routes
	api.GET("/asteroides/:id", controllers.GetAsteroid)
	api.GET("/asteroides", controllers.GetAsteroids)
	api.POST("/asteroides", controllers.CreateAsteroid)
	api.PUT("/asteroides/:id", controllers.UpdateAsteroid)
	api.DELETE("/asteroides/:id", controllers.DeleteAsteroid)
}

func serverHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("x-version", "Test/v1.0")
		return next(c)
	}
}
