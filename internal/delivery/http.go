package delivery

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type EchoHttpService struct {
	Echo *echo.Echo
}

func (e *EchoHttpService) Start(address string) {
	e.Echo.Logger.Fatal(e.Echo.Start(address))
}

func NewEchoHttpService() *EchoHttpService {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
	}))

	return &EchoHttpService{
		Echo: e,
	}
}
