package go_database

import (
	"database/sql"
	"fmt"
	"time"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/go_database?parseTime=true")
	if err != nil {
		panic(err)
	}

	fmt.Println("DB Connected!")
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)
	return db
}
