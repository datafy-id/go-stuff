package main

import (
	"github.com/labstack/echo/v4"
	"net/http"

	"github.com/datafy-id/go-stuff/hello/pkg/greetings"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, greetings.Hello())
	})
	e.Logger.Fatal(e.Start(":8000"))
}
