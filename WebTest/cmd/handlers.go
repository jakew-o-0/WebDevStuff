package main

import (
	"webtest/templates"

	"github.com/labstack/echo/v4"
)

func indexPageGet(c echo.Context) error {
	return renderPage(templates.IndexPage(), c)
}

func loginPageGet(c echo.Context) error {
	return renderPage(templates.LoginPage(), c)
}

func signupPageGet(c echo.Context) error {
	return renderPage(templates.SignupPage(), c)
}
