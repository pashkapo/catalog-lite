package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
	"github.com/pashkapo/catalog-lite/config"
	database "github.com/pashkapo/catalog-lite/db"
	"github.com/pashkapo/catalog-lite/handler"
	"net/http"
	"time"
)

func main() {
	c := config.New()
	db, err := database.New(c)
	if err != nil {
		log.Fatal(err)
	}
	h := handler.New(db)

	e := echo.New()
	e.Logger.SetLevel(log.ERROR)
	//e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	g := e.Group("/api")
	g.GET("/ping", h.Ping)
	g.GET("/firms", h.GetFirms)
	g.GET("/firms/:id", h.GetFirmById)
	g.GET("/buildings", h.GetBuildings)

	s := &http.Server{
		Handler:      e,
		Addr:         ":" + c.AppPort,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	e.Logger.Fatal(e.StartServer(s))
}
