package model

// UsersTable ...
const UsersTable = "users"

// User - пользователи
type User struct {
	ID    int32  `db:"id"`
	FIO   string `db:"fio"`
	Phone string `db:"phone"`
	Age   int32  `db:"age"`
	Level int32  `db:"level"`
}
