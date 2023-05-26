package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Employee struct {
	ID           uint
	Name         string
	Address      string
	Age          uint8
	Birthdate    string
	LevelStaff   string
	LevelSpv     string
	LevelManager string
}

func main() {
	dsn := "host=localhost  user=aghniam password=mudahkanLah1602 dbname=test_db_camp port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.Create(&Employee{ID: 123, Name: "Aghnia", Address: "Bandung", Age: 21, Birthdate: "2002-04-16", LevelStaff: "staff"})

		if err := db.Error; err != nil {
			fmt.Println(err)
		}

		fmt.Println("Insert Success")
	
}