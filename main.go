package main

import (
	"encoding/json"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
	"github.com/pashkapo/catalog-lite/core"
	"github.com/pashkapo/catalog-lite/db"
	"github.com/pashkapo/catalog-lite/models"
	"net/http"
	"strconv"
)

func main() {
	config := core.NewConfig()

	database, err := db.New(config)
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	e.Use(middleware.Recover())

	g := e.Group("/api")

	g.GET("/ping", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	// @toDo вынести
	g.GET("/firms", func(c echo.Context) error {
		page, _ := strconv.Atoi(c.QueryParam("page"))
		count, _ := strconv.Atoi(c.QueryParam("count"))
		filter := models.FirmFilter{BuildingId: 0}
		err := json.Unmarshal([]byte(c.QueryParam("filter")), &filter)

		firms, err := database.GetFirms(page, count, filter)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, models.Error{Message: err.Error()})
		}
		return c.JSON(http.StatusOK, firms)
	})

	g.GET("/firms/:id", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))

		firm, err := database.GetFirmById(uint(id))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, models.Error{Message: err.Error()})
		}
		return c.JSON(http.StatusOK, firm)
	})

	g.GET("/buildings", func(c echo.Context) error {
		page, _ := strconv.Atoi(c.QueryParam("page"))
		count, _ := strconv.Atoi(c.QueryParam("count"))

		firms, err := database.GetBuildings(page, count)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, models.Error{Message: err.Error()})
		}
		return c.JSON(http.StatusOK, firms)
	})

	e.Logger.Fatal(e.Start(":" + config.AppPort))
}
