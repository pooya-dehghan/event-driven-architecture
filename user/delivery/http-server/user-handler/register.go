package userhandler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h Handler) register(c echo.Context) error {
	resp, err := h.userSvc.AddExpenseHandler()
	if err != nil {
		return fmt.Errorf("expense did not add")
	}

	return c.JSON(http.StatusOK, resp)
}
