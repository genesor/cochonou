package http

import (
	"net/http"

	"github.com/labstack/echo"
)

type HelloHandler struct{}

func NewHelloHandler() *HelloHandler {
	return &HelloHandler{}
}

func (h *HelloHandler) HandleHello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
