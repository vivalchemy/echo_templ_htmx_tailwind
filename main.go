package main

import (
	"errors"
	"os"

	"github.com/a-h/templ"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/vivalchemy/echo_templ_htmx_tailwind/components"
	"github.com/vivalchemy/echo_templ_htmx_tailwind/routes"
	"net/http"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		errors.New("Error loading the .env file")
	}
	e := echo.New()
	e.Static("/build", "build")
	e.Use(middleware.Logger())
	// e.Use(middleware.Recover())
	e.GET("/", HomeHandler)
	e.GET("/new", NewHandler)

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}

func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)

	if err := t.Render(ctx.Request().Context(), buf); err != nil {
		return err
	}

	return ctx.HTML(statusCode, buf.String())
}

func HomeHandler(c echo.Context) error {
	return Render(c, http.StatusOK, routes.Index("Title", components.Time()))
}

func NewHandler(c echo.Context) error {
	return Render(c, http.StatusOK, components.Time())
}

// package main
//
// import (
// 	"context"
// 	"github.com/vivalchemy/echo_templ_htmx_tailwind/components"
// 	"os"
// )
//
// func main() {
// 	component := components.Hello("John")
// 	component.Render(context.Background(), os.Stdout)
// }
