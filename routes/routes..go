package routes

import (
	"converter-iso8583/controllers"

	"github.com/labstack/echo/v4"
)

func Routes() *echo.Echo {
	e := echo.New()

	e.POST("/check", controllers.EchoTestControllers)

	return e
}
