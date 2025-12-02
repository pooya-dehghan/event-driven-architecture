package userhandler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pooya/params"
)

func (h Handler) register(c echo.Context) error {
	resp, err := h.userSvc.Register()

	if err != nil {
		return fmt.Errorf("expense did not add")
	}

	return c.JSON(http.StatusOK, resp)
}

func (h Handler) currencyRequest(c echo.Context) error {
	var req params.CurrencyRequestParams

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	resp, err := h.userSvc.AddCurrencyRequestHandler(req)

	if err != nil {
		return fmt.Errorf("expense did not add")
	}

	return c.JSON(http.StatusOK, resp)
}
