package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterHelloWorldHandler(e *echo.Echo) {
	e.GET("/api/helloworld", HelloWorldHandler)
}

func HelloWorldHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
