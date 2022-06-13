package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.GET("/health", func(c echo.Context) error {
		fmt.Print("test")
		return c.NoContent(http.StatusOK)
	})

	e.Start(":8080")
}
