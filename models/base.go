package models

type Base struct {
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
	Deleted   bool   `db:"deleted"`
}
