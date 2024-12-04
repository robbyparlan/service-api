package seeders

import (
	"log"
	"service-api/models"
	"service-api/utils"

	"gorm.io/gorm"
)

/*
Initialize Database
@param db *gorm.DB
*/
func Seed(db *gorm.DB) {
	// Cek jika tabel users sudah memiliki data
	if !isTableEmpty(db, &models.Users{}) {
		log.Println("--------------------- : Table users is not empty. Skipping truncate.")
	} else {
		truncateTable(db)
		seedUsers(db)
	}
}

/*
Truncate Table Before Seed
@param db *gorm.DB
*/
func truncateTable(db *gorm.DB) {
	// Deactivate foreign key check
	db.Exec("SET session_replication_role = 'replica';")

	// Truncate table
	db.Exec("TRUNCATE TABLE users RESTART IDENTITY CASCADE;")

	// Activate foreign key check
	db.Exec("SET session_replication_role = 'origin';")
}

/*
Check if Table is Empty
@param db *gorm.DB, model interface{}
*/
func isTableEmpty(db *gorm.DB, model interface{}) bool {
	var count int64
	if err := db.Model(model).Count(&count).Error; err != nil {
		log.Fatalf("Failed to check table records: %v", err)
	}
	return count == 0
}

/*
Seed Users
@param db *gorm.DB
*/
func seedUsers(db *gorm.DB) {
	users := []models.Users{
		{Username: "admin", Password: utils.HashedPassword("admin"), Role: "admin"},
		{Username: "user", Password: utils.HashedPassword("user"), Role: "user"},
	}

	for _, user := range users {
		err := db.Create(&user).Error
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Println("Seed Users Done")
}
