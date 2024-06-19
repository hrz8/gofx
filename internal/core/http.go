package core

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

type Router struct {
	Mux *echo.Echo
}

func (r *Router) Pre(middleware ...echo.MiddlewareFunc) {
	r.Mux.Pre(middleware...)
}

func (r *Router) Use(middleware ...echo.MiddlewareFunc) {
	r.Mux.Use(middleware...)
}

func (r *Router) Any(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) []*echo.Route {
	return r.Mux.Any(path, h, m...)
}

func (r *Router) GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return r.Mux.GET(path, h, m...)
}

func (r *Router) POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return r.Mux.POST(path, h, m...)
}

func (r *Router) PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return r.Mux.PATCH(path, h, m...)
}

func (r *Router) PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return r.Mux.PUT(path, h, m...)
}

func (r *Router) DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return r.Mux.DELETE(path, h, m...)
}

func (r *Router) RouteNotFound(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return r.Mux.RouteNotFound(path, h, m...)
}

func (r *Router) Match(methods []string, path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) []*echo.Route {
	return r.Mux.Match(methods, path, h, m...)
}

func NewRouter() *Router {
	e := echo.New()
	return &Router{Mux: e}
}

func NewHTTPServer(lc fx.Lifecycle, logger *Logger, router *Router, cfg *Config) *http.Server {
	logger.Info(
		"registering http server",
		slog.String("port", cfg.AppPort),
	)
	srv := &http.Server{
		Addr:    ":" + cfg.AppPort,
		Handler: router.Mux,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				err := srv.ListenAndServe()
				if err != nil && !errors.Is(err, http.ErrServerClosed) {
					logger.Error("error starting http server", "err", err)
					os.Exit(1)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			toCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
			defer cancel()

			if err := srv.Shutdown(toCtx); err != nil {
				logger.Error(
					"can't shutdown http server gracefully",
					slog.String("err", err.Error()),
				)
				return err
			}

			logger.Info("gracefully shutdown http server")
			return nil
		},
	})

	return srv
}
