package repository

import (
	"fmt"
	"log"
	"zargram/configs"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDBConnection(cfg configs.DatabaseConnConfig) *gorm.DB {
	dsn := cfg.GetDSN()

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("Error connecting database. Error is %v", err.Error())
		panic(err.Error())
	}

	log.Printf("Connection success host:%s port:%s", cfg.Host, cfg.Port)

	return db
}

func Close(db *gorm.DB) {
	conn, err := db.DB()
	if err != nil {
		log.Fatalln(err.Error())
	}
	err = conn.Close()
}
