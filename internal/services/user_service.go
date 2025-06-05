package services

import (
	"database/sql"
	"errors"


	"github.com/kskr24/workspacehub/internal/db"
	model "github.com/kskr24/workspacehub/internal/models"
)

func GetUserByID(id int64) (*model.User, error) {
	var user model.User
	row := db.DB.QueryRow("SELECT id, name, email, created_at FROM users WHERE id = $1", id)

	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func UpdateUser(user *model.User) error {
	_, err := db.DB.Exec(`
	UPDATE users SET name = $1, email = $2 where id = $3`, user.Name, user.Email, user.ID)
	return err
}
