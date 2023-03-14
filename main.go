package main

import (
	"log"

	"github.com/Emmrys-Jay/altschool-sms/internal/config"
	mdb "github.com/Emmrys-Jay/altschool-sms/pkg/repository/storage/mongo"
	"github.com/Emmrys-Jay/altschool-sms/utility"
	"github.com/go-playground/validator/v10"

	// "github.com/Emmrys-Jay/altschool-sms/pkg/repository/storage/redis"
	"github.com/Emmrys-Jay/altschool-sms/pkg/router"
)

func init() {
	config.Setup()
	// redis.SetupRedis() uncomment when you need redis
	mdb.ConnectToDB()
}

func main() {
	//Load config
	logger := utility.NewLogger()
	getConfig := config.GetConfig()
	validatorRef := validator.New()
	r := router.Setup(validatorRef, logger)

	logger.Info("Server is starting at 127.0.0.1:%s", getConfig.Server.Port)
	log.Fatal(r.Run(":" + getConfig.Server.Port))
}
