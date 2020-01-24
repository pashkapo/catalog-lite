package handler

import (
	"github.com/labstack/echo"
	"github.com/pashkapo/catalog-lite/db"
	"net/http"
)

type Handler struct {
	DB *db.Database
}

func New(db *db.Database) *Handler {
	return &Handler{DB: db}
}

func (h *Handler) Ping(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}
