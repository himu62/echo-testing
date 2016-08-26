package route

import (
	"github.com/himu62/echo-testing/controller"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
)

func Init() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	controller.AddRoutes(e)

	return e
}

func Run(e *echo.Echo) {
	e.Run(standard.New(":8888"))
}
