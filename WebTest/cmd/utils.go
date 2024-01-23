package main

import (
	"webtest/templates"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func render(t templ.Component, c echo.Context) error {
	return t.Render(c.Request().Context(), c.Response().Writer)
}

func renderPage(t templ.Component, c echo.Context) error {
	return render(templates.BaseHTML(t), c)
}
