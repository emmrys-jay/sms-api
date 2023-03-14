package mysql

import (
	"fmt"
	"github.com/Emmrys-Jay/altschool-sms/internal/config"
	"github.com/Emmrys-Jay/altschool-sms/utility"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var (
	db *gorm.DB
)

func DB() *gorm.DB {
	return db
}

// dsn forms the connection string
func dsn() string {
	config := config.GetConfig()
	username := config.DBUsername
	password := config.DBPassword
	host := config.DBHost
	port := config.DBPort
	dbname := config.DBName
	return username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
}

// ConnectToDB connects to the mysql database
func ConnectToDB() *gorm.DB {
	logger := utility.NewLogger()
	database, err := gorm.Open(mysql.Open(dsn()), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	// IF EVERYTHING IS OKAY, THEN CONNECT
	fmt.Println("MYSQL CONNECTION ESTABLISHED")
	logger.Info("MYSQL CONNECTION ESTABLISHED")

	db = database
	return db
}
