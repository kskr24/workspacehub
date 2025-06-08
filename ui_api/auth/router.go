package auth

import (
	"github.com/kskr24/sajha/ui_api/web"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func SetupRoutes(e *echo.Echo, l *zap.Logger) {
	//e.POST("/v1/auth/signup", web.WrapPublicRoute(signup, l))
	e.POST("/v1/auth/login", web.WrapPublicRoute(login, l))
}
