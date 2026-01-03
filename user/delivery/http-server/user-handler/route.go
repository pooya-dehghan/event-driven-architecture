package userhandler

import "github.com/labstack/echo/v4"

func (h Handler) SetUserRoute(e *echo.Echo) {
	userGroup := e.Group("/users")

	userGroup.POST("/expense", h.expensePublisher)
}
