package model

import "time"

// TemplateTableName - название таблицы в бд
const TemplateTableName = "template"

// Template - операция tsys
type Template struct {
	Name        string    `db:"name"`
	Description *string   `db:"description"`
	Status      int32     `db:"status"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}
