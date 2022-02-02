package classes

import "database/sql"

type Message struct {
	ID       int
	Body     string
	SendTime string
	RecvTime sql.NullString
	ReadTime sql.NullString
	UserID   int
	ChatID   int
}
