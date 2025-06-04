package model

type User struct {
	ID        int64  `db:"id" json:"id"`
	Name      string `db:"name" json:"name"`
	Email     string `db:"email" json:"email"`
	Password  string `db:"password" json:"-"`
	CreatedAt string `db:"created_at" json:"created_at"`
}
