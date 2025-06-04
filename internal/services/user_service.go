package services

import (
	"database/sql"
	"errors"

	"github.com/kskr24/workspacehub/internal/db"
	"github.com/kskr24/workspacehub/internal/model"
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
