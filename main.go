package main

import (
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "Achievement/backend/app/route"
)

func main() {
		e := echo.New()
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())
    e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
      AllowOrigins: []string{
        "http://localhost:8080",
      },
      AllowCredentials: true,
    }))
    route.Route(e)
    e.Logger.Fatal(e.Start(":8888"))
  }