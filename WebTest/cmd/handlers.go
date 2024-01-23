package main

import (
	"webtest/templates"

	"github.com/labstack/echo/v4"
)

func indexPageGet(c echo.Context) error {
	return renderPage(templates.IndexPage(), c)
}
