package auth

import "context"

func FetchSessionByToken(ctx context.Context, token string) (*Session, error) {
	return &Session{}, nil
}

func FetchUserByEmail(ctx context.Context, email string) (*User, error) {
	return &User{}, nil
}
