package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Secure())


	e.GET("/", indexPageGet)
	e.Static("/static", "./static")

	e.GET("/login", loginPageGet)
	e.POST("/login", loginPagePost)
	e.GET("/signup", signupPageGet)

	e.GET("/workouts", workoutsPageGet)
	e.GET("/diets", dietsPageGet)


	e.Logger.Fatal(e.Start(":3000"))
}
