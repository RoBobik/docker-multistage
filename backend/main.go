package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Quote struct {
	Text   string `json:"text"`
	Author string `json:"author"`
}

var quotes = []Quote{
	{"There are only two ways to live your life. One is as though nothing is a miracle. The other is as though everything is a miracle.", "Albert Einstein"},
	{"I have not failed. I've just found 10,000 ways that won't work.", "Thomas A. Edison"},
	{"A day without sunshine is like, you know, night.", "Steve Martin"},
	{"Some people never go crazy. What truly horrible lives they must lead.", "Charles Bukowski"},
	{"I may not have gone where I intended to go, but I think I have ended up where I needed to be.", "Douglas Adams"},
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Serve frontend
	e.Static("/", "public")

	// API routes
	e.GET("/api/quotes/random", func(c echo.Context) error {
		time.Sleep(time.Second) // Artificial delay

		return c.JSON(http.StatusOK, quotes[rand.Intn(len(quotes))])
	})

	e.Logger.Fatal(e.Start(":8080"))
}
