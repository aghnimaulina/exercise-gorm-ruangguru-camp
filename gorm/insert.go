package main

import (
	"database/sql"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID           uint
	Name         string
	Email        string
	Age          uint8
	Birthday     string
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func main() {
	dsn := "host=localhost  user=aghniam password=mudahkanLah1602 dbname=exercise_db port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")

		//create
		db.Create(&User{Name: "Aghnia", Email: "test@mail.com", Age: 21, Birthday: "2002-04-16", MemberNumber: sql.NullString{String: "123", Valid: true}, ActivatedAt: sql.NullTime{Time: time.Now(), Valid: true}})

		if err := db.Error; err != nil {
			fmt.Println(err)
		}

		fmt.Println("Migration Success")
	}
}
