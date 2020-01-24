package handler

import (
	"encoding/json"
	"github.com/labstack/echo"
	"github.com/pashkapo/catalog-lite/model"
	"net/http"
	"strconv"
)

func (h *Handler) GetFirms(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	count, _ := strconv.Atoi(c.QueryParam("count"))
	filter := new(model.FirmFilter)
	err := json.Unmarshal([]byte(c.QueryParam("filter")), &filter)

	firms, err := h.DB.GetFirms(page, count, filter)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.Error{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, firms)
}

func (h *Handler) GetFirmById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	firm, err := h.DB.GetFirmById(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.Error{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, firm)
}
