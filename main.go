package main

import (
	"github.com/Emmrys-Jay/altschool-sms/pkg/repository/storage/mysql"
	"log"

	"github.com/Emmrys-Jay/altschool-sms/internal/config"
	"github.com/Emmrys-Jay/altschool-sms/utility"
	"github.com/go-playground/validator/v10"

	// "github.com/Emmrys-Jay/altschool-sms/pkg/repository/storage/redis"
	"github.com/Emmrys-Jay/altschool-sms/pkg/router"
)

func init() {
	config.Setup()
	// redis.SetupRedis() uncomment when you need redis
	mysql.ConnectToDB()
}

//	@title			School Management System
//	@version		1.0
//	@description	This is a School Management System API.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:3000
//	@BasePath	/api/v1

//	@securityDefinitions.basic	BasicAuth

//	@externalDocs.description	OpenAPI
//	@externalDocs.url			https://swagger.io/resources/open-api/

func main() {
	logger := utility.NewLogger()

	//Migrate DB
	if err := mysql.MigrateDB(logger); err != nil {
		log.Fatal(err)
	}

	//Load config
	getConfig := config.GetConfig()
	validatorRef := validator.New()
	r := router.Setup(validatorRef, logger)

	logger.Info("Server is starting at 127.0.0.1:%s", getConfig.Port)
	log.Fatal(r.Run(":" + getConfig.Port))
}
