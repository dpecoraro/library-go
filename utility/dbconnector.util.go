package utility

import (
	"fmt"
	"os"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type DbConector struct {
}

const (

)



func (connector *DbConector) Connect() *gorm.DB{
	server := os.Getenv("SQL_SERVER")
	port := os.Getenv("SQL_PORT")
	user := os.Getenv("SQL_USER")
	password := os.Getenv("SQL_PASSWORD")
	database := os.Getenv("SQL_DATABASE")

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s;",
		server, user, password, port, database)
	
	db, err := gorm.Open(sqlserver.Open(connString))

	if err != nil {
		fmt.Println(err.Error())
	}
	return db;
}