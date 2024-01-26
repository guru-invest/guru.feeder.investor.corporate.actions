package repository

import (
	"fmt"
	"log"
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

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Discard})
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

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Discard})
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
