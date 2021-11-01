package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Name struct {
	gorm.Model
	Firstname  string
	Secondname string
}

func main() {

	db, err := gorm.Open(sqlite.Open("student.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect")

	}

	db.AutoMigrate(&Name{})

	db.Create(&Name{Firstname: "Alex", Secondname: "Mumo"})

	var name Name
	db.Delete(&name, 1)

}
