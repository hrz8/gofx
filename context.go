package gofx

import (
	"context"
	"io"
	"io/fs"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Component interface {
	Render(ctx context.Context, w io.Writer) error
}

type Context struct {
	echo.Context
}

func (c *Context) IsHtmx() bool {
	return c.Request().Header.Get("Hx-Request") != "true"
}

func (c *Context) RenderView(code int, component Component) error {
	w := c.Response().Writer
	w.WriteHeader(code)
	return component.Render(c.Request().Context(), w)
}

func (c *Context) ServeStatic(path string, file fs.File) error {
	ff, ok := file.(io.ReadSeeker)
	if !ok {
		return c.String(http.StatusOK, "not found")
	}
	fileInfo, _ := file.Stat()
	http.ServeContent(c.Response().Writer, c.Request(), path, fileInfo.ModTime(), ff)
	return nil
}
