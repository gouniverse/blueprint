package middlewares

import (
	"github.com/gouniverse/dashboard"
	"github.com/gouniverse/router"
)

func NewThemeMiddleware() router.Middleware {
	m := router.Middleware{
		Name:    "Theme Middleware",
		Handler: dashboard.ThemeMiddleware,
	}
	return m
}
