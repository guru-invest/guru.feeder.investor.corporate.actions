package repository

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/guru-invest/guru.feeder.investor.corporate.actions/src/crossCutting/options"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DatabaseConnection struct {
	_databaseConnection *gorm.DB
}

func (db *DatabaseConnection) connect() {

	DATABASE := options.OPTIONS.DATABASE
	dsn := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s",
		DATABASE.Username,
		DATABASE.Password,
		DATABASE.Database,
		DATABASE.Url,
		DATABASE.Port)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Millisecond * 300, // Slow SQL threshold
			LogLevel:                  logger.Error,           // Log level
			IgnoreRecordNotFoundError: false,                  // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,                   // Disable color
		},
	)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 newLogger,
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	sqlDB, err := database.DB()
	if err != nil {
		log.Fatal(err.Error())
	}
	sqlDB.SetMaxOpenConns(50)                 // maximo de nova conexao por pool de conexao
	sqlDB.SetMaxIdleConns(25)                 // maximo de conexão inativa aguardando reuso
	sqlDB.SetConnMaxLifetime(3 * time.Minute) // tempo maximo para expirar uma conexao

	db._databaseConnection = database

}

func (db *DatabaseConnection) connectStateLess() {

	DATABASE := options.OPTIONS.DATABASE
	dsn := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s",
		DATABASE.Username,
		DATABASE.Password,
		DATABASE.Database,
		DATABASE.Url,
		DATABASE.Port)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Millisecond * 300, // Slow SQL threshold
			LogLevel:                  logger.Error,           // Log level
			IgnoreRecordNotFoundError: false,                  // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,                   // Disable color
		},
	)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 newLogger,
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	db._databaseConnection = database

}

func (db *DatabaseConnection) disconnect() {
	sqlDB, err := db._databaseConnection.DB()
	if err != nil {
		log.Fatal(err.Error())
	}
	sqlDB.Close()
}
