package auth

import "golang.org/x/crypto/bcrypt"

type UserRole string

const (
	RoleAdmin UserRole = "admin"
	RoleUser  UserRole = "user"
)

type User struct {
	ID                 int64
	Name               string
	Email              string
	Password           string
	Verified           bool
	Role               UserRole
	DefaultWorkspaceID int64
	Created            int64
	Updated            int64
}

func (u *User) verifyPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func (u *User) hashPassword() error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(bytes)
	return nil
}

func userRoleFromString(role string) UserRole {
	switch role {
	case "admin":
		return RoleAdmin
	case "user":
		return RoleUser
	default:
		return RoleUser
	}
}
