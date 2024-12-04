package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

/*
* Menghubungkan ke database
 */
func NewDB(cfg DBConfig) (*gorm.DB, error) {
	log.Println("--------------------- : Database connecting")
	// Menghubungkan ke database
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=Asia/Jakarta",
		cfg.DBHost,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBPort,
		cfg.DBSSL,
	)

	// Menambahkan logger GORM dengan level Info untuk melihat query SQL
	gormLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // output log ke stdout
		logger.Config{
			LogLevel: logger.Info, // Menampilkan query, waktu eksekusi, dan parameter
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("--------------------- : Database connected")

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	sqlDB.SetMaxOpenConns(cfg.DBMaxOpenConn)
	sqlDB.SetMaxIdleConns(cfg.DBIdleConn)
	sqlDB.SetConnMaxLifetime(time.Minute * time.Duration(cfg.DBConnMaxLifeTime))

	log.Println("--------------------- : Database connection initialized and pool configured.")
	return db, nil
}
