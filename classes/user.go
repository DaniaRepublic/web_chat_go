package classes

import (
	"database/sql"
)

// structs
type User struct {
	ID         int
	Username   string
	PhoneNum   string
	Bio        sql.NullString
	FirstName  sql.NullString
	LastName   sql.NullString
	CreateTime string
}

// methods
