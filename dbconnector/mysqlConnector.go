package dbconnector

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/DaniaRepublic/commonSpaceGo/classes"

	"github.com/go-sql-driver/mysql"
)

type MYSQLConn struct {
	DB *sql.DB
}

func (conn *MYSQLConn) Connect() {
	cfg := mysql.Config{
		User:   os.Getenv("MYSQLuser"),
		Passwd: os.Getenv("MYSQLpass"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "common_space",
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(fmt.Errorf("error opening MYSQL: %v", err))
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(fmt.Errorf("error in MYSQL ping: %v", pingErr))
	}

	conn.DB = db
}
