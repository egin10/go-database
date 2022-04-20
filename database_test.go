package go_database

import (
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestDatabase(t *testing.T) {

}

func TestOpenConnection(t *testing.T) {
	db := GetConnection()
	defer db.Close()
}
