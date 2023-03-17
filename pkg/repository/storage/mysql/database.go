package mysql

import (
	"context"
	"fmt"
	"github.com/Emmrys-Jay/altschool-sms/internal/config"
	"github.com/Emmrys-Jay/altschool-sms/internal/model"
	"github.com/Emmrys-Jay/altschool-sms/pkg/repository/storage"
	"github.com/Emmrys-Jay/altschool-sms/utility"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

type MySQL struct {
	db *gorm.DB
}

var (
	db                  *gorm.DB
	generalQueryTimeout = 60 * time.Second
)

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

func MigrateDB(logger *utility.Logger) error {
	err := db.AutoMigrate(&model.Student{}, &model.Course{})
	if err != nil {
		return err
	}

	logger.Info("DATABASE MIGRATION SUCCESSFUL")
	return nil
}

func GetDB() storage.Repository {
	return &MySQL{db: db}
}

// DBWithTimeout returns a database with timeout, and the context's cancel func
func (m *MySQL) DBWithTimeout(ctx context.Context) (*gorm.DB, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(ctx, generalQueryTimeout)
	return m.db.WithContext(ctx), cancel
}

// dsn forms the connection string
func dsn() string {
	conf := config.GetConfig()
	username := conf.DBUsername
	password := conf.DBPassword
	host := conf.DBHost
	port := conf.DBPort
	dbname := conf.DBName
	return username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
}
