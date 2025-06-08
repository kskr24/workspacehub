package auth

import (
	"context"
	"time"

	"github.com/bluele/go-timecop"
)

func Login(ctx context.Context, email, password string, remember bool, ip, ua string) (*Bundle, error) {
	user, err := FetchUserByEmail(ctx, email)

	if err != nil {
		return nil, err
	}

	if !user.verifyPassword(password) {
		return nil, ErrInvalidPassword
	}

	// TODO(get a list of workspaces)

	now := timecop.Now()
	expires := now.Add(30 * 24 * time.Hour)

	if remember {
		expires = now.Add(365 * 24 * time.Hour)
	}

	session := &Session{
		UA:      ua,
		IP:      ip,
		UserID:  user.ID,
		Token:   GenerateToken(user),
		Expires: expires.Unix(),
		Created: now.Unix(),
		Updated: now.Unix(),
	}

	// TODO (Insert Session into table)
	return &Bundle{
		User:    user,
		Session: session,
	}, nil

}
