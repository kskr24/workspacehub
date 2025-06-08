package web

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/bluele/go-timecop"
	"github.com/kskr24/sajha/domains/auth"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type Context struct {
	echo.Context

	L           *zap.Logger
	Session     *auth.Session
	User        *auth.User
	WorkspaceID int64
}

// TODO(crate a random string lib)
func (c Context) InternalError(err error) error {
	ref := fmt.Sprintf("#%s", "s")

	resp := map[string]any{
		"title":  fmt.Sprintf("Internal Error Reference: %s", ref),
		"detail": "Please try again. If the issue persists connect with us with the reference number",
		"type":   "alert",
	}

	c.L.Error("internal error", zap.String("ref", ref), zap.Error(err))
	return c.JSON(http.StatusInternalServerError, resp)
}

func (c Context) Unauthorized() error {
	c.SetCookie(&http.Cookie{
		Name:     "session_token",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	})

	return c.JSON(http.StatusUnauthorized, map[string]string{
		"error": "Unauthorized",
	})
}

type HandlerFunc func(ctx Context) error

func WrapPublicRoute(h HandlerFunc, l *zap.Logger) echo.HandlerFunc {
	return public(h, l)
}

func public(h HandlerFunc, l *zap.Logger) echo.HandlerFunc {
	return func(c echo.Context) error {
		rqid := c.Response().Header().Get(echo.HeaderXRequestID)

		ctx := Context{
			Context: c,
			L:       l.With(zap.String("request_id", rqid)),
		}

		return h(ctx)
	}
}

func authenticated(h HandlerFunc, l *zap.Logger) echo.HandlerFunc {
	return func(c echo.Context) error {
		rqid := c.Response().Header().Get(echo.HeaderXRequestID)

		ctx := Context{
			Context: c,
			L:       l.With(zap.String("request_id", rqid)),
		}

		currentTime := timecop.Now().Unix()

		tkCookie, err := c.Request().Cookie("session_token")
		if err != nil && errors.Is(err, http.ErrNoCookie) {
			return ctx.Unauthorized()
		}

		rctx := c.Request().Context()
		session, err := auth.FetchSessionByToken(rctx, tkCookie.Value)

		if err != nil || session.Expires < currentTime {
			return ctx.Unauthorized()
		}

		ctx.Session = session

		return h(ctx)
	}

}
