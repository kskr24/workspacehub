package auth

import (
	"errors"
	"net/http"
	"net/mail"
	"strings"

	"github.com/kskr24/sajha/domains/auth"
	"github.com/kskr24/sajha/ui_api/web"
)

func login(ctx web.Context) error {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Remember bool   `json:"remember"`
	}

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, map[string]any{
			"error": err.Error(),
		})
	}

	req.Email = strings.TrimSpace(req.Email)

	_, err := mail.ParseAddress(req.Email)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]any{
			"email": "User not found",
		})
	}

	bundle, err := auth.Login(ctx.Request().Context(), req.Email, req.Password, req.Remember, ctx.RealIP(), ctx.Request().Header.Get("User-Agent"))

	if err != nil {
		if errors.Is(err, auth.ErrUserNotFound) {
			return ctx.JSON(http.StatusBadRequest, map[string]any{
				"email": "User not found with this email",
			})
		}

		if errors.Is(err, auth.ErrInvalidPassword) {
			return ctx.JSON(http.StatusBadRequest, map[string]any{
				"password": "Password is incorrect",
			})
		}

		return ctx.InternalError(err)
	}

	cookie := &http.Cookie{
		Name:     "session_token",
		Value:    bundle.Session.Token,
		Path:     "/",
		MaxAge:   int(bundle.Session.Age().Seconds()),
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	}
	ctx.SetCookie(cookie)

	return ctx.JSON(http.StatusOK, req)
}
