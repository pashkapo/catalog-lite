package handler

import (
	"github.com/labstack/echo"
	"github.com/pashkapo/catalog-lite/model"
	"net/http"
	"strconv"
)

func (h *Handler) GetBuildings(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	count, _ := strconv.Atoi(c.QueryParam("count"))

	firms, err := h.DB.GetBuildings(page, count)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.Error{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, firms)
}
