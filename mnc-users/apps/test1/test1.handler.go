package test1

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Test1Handler  struct {
	Service *Test1Service
}

func NewTest1Handler(service *Test1Service) *Test1Handler {
	return &Test1Handler{Service: service}
}

func (h *Test1Handler) FindMatchString(c echo.Context) error {
	result := h.Service.FindMatchString(4, []string{"abcd", "acbd", "aaab", "acbd"})
	return c.JSON(http.StatusOK, result)
}

func (h *Test1Handler ) CalculateChange(c echo.Context) error {
	result := h.Service.CalculateChange(700649, 800000)
	return c.JSON(http.StatusOK,result)
}

func (h *Test1Handler ) IsValidBracketSequence(c echo.Context) error {
	result := h.Service.IsValidBracketSequence("{{[<>[{{}}]]}}")
	return c.JSON(http.StatusOK, result)
}

func (h *Test1Handler ) CheckLeave(c echo.Context) error {
	result, message := h.Service.CheckLeave(7, "2021-05-01", "2021-07-05", 1)
	return c.JSON(http.StatusOK, echo.Map{"success": result, "message": message})
}