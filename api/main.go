package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Serve() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "api v1.0")
	})

	e.GET("/users/:id", func(c echo.Context) error {
		id := c.Param("id")
		return c.String(http.StatusOK, fmt.Sprintf("User ID: %v", id))
	})

	e.Start(":8080")
}
