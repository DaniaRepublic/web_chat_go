package classes

import "database/sql"

type Chat struct {
	ID          int
	Title       sql.NullString
	Description sql.NullString
	CreateTime  string
	GroupChat   bool
}
