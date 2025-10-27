package httpserver

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s Server) addExpense(c echo.Context) error {
	resp, err := s.userService.AddExpenseHandler()
	if err != nil {
		return fmt.Errorf("expense did not add")
	}

	return c.JSON(http.StatusOK, resp)

}
