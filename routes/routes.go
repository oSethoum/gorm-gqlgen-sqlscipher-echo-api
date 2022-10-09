package routes

import (
	"gormql/handlers"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo) {
	e.Any("/", handlers.GqlPlaygroundHandler)
	e.Any("/query", handlers.GqlHandler)
}
