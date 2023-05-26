package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Department struct {
	ID             uint
	NameDepartment string
}

func main() {
	dsn := "host=localhost  user=aghniam password=mudahkanLah1602 dbname=test_db_camp port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	//create
	db.Create(&Department{ID: 05, NameDepartment: "HR"})

	if err := db.Error; err != nil {
		fmt.Println(err)
	}

	fmt.Println("Insert Success")
}
