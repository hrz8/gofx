package gofx

import (
	"io"
	"io/fs"
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

type Context struct {
	echo.Context
}

func (c *Context) RenderView(component templ.Component) error {
	return component.Render(c.Request().Context(), c.Response().Writer)
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
