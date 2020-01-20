package main

import (
	"database/sql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
	"gitlab.com/pashkapo/gis_catalog/models"
	"net/http"
	"os"
)

func init() {
	_ = os.Setenv("PORT", "3000")
	_ = os.Setenv("DATABASE_URL", "postgresql://postgres:postgres@0.0.0.0:5432/catalog-lite?sslmode=disable")
}

func main() {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("could not ping DB... %v", err)
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	g := e.Group("/api")
	g.GET("/ping", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})
	g.GET("/firms", getFirms)

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}

func getFirms(c echo.Context) error {
	firms := &models.Firms{
		models.Firm{
			Id:   1,
			Name: "1gis",
		},
		models.Firm{
			Id:   2,
			Name: "2gis",
		},
	}
	return c.JSON(http.StatusOK, firms)
}
