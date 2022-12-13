package providers

import (
	"log"
	"time"

	"github.com/guru-invest/guru.feeder.investor.corporate.actions.oms/src/crossCutting/options"
	database_connector "github.com/guru-invest/guru.framework/src/infrastructure/database-connector"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DatabaseConnection struct {
	DatabaseConnection *gorm.DB
}

func (db *DatabaseConnection) Connect() {

	DATABASE := options.OPTIONS.DATABASE
	sqlDB, err := database_connector.DatabaseConnector{
		Port:               DATABASE.Port,
		URL:                DATABASE.Url,
		Username:           DATABASE.Username,
		Password:           DATABASE.Password,
		Database:           DATABASE.Database,
		SetConnMaxLifetime: DATABASE.ConnMaxLifetime * time.Second,
	}.ConnectForServerless()
	if err != nil {
		log.Fatal(err.Error())
	}

	gormDB, err := gorm.Open(
		postgres.New(postgres.Config{Conn: sqlDB}),
		&gorm.Config{Logger: logger.Discard},
	)
	if err != nil {
		log.Fatal(err.Error())
	}
	db.DatabaseConnection = gormDB

}

// func (db *DatabaseConnection) connectStateLess() {

// 	DATABASE := options.OPTIONS.DATABASE
// 	dsn := fmt.Sprintf(
// 		"user=%s password=%s dbname=%s host=%s port=%s",
// 		DATABASE.Username,
// 		DATABASE.Password,
// 		DATABASE.Database,
// 		DATABASE.Url,
// 		DATABASE.Port)

// 	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
// 		SkipDefaultTransaction: true,
// 		Logger:                 logger.Default,
// 	})
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}

// 	db._databaseConnection = database

// }

// func (db *DatabaseConnection) disconnect() {
// 	sqlDB, err := db._databaseConnection.DB()
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}
// 	sqlDB.Close()
// }
