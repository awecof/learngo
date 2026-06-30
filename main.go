package main

import (
	"learngo/scrapper"
	"strings"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func handleHome(c *echo.Context) error {
	return c.File("home.html")
}

func handleScrape(c *echo.Context) error {
	term := strings.ToLower(c.FormValue("term"))
	scrapper.Scrape(term)
	return nil
}

func main() {
	scrapper.Scrape("python")
	e := echo.New()
	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())
	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)
	if err := e.Start(":1323"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
