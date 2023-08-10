package models

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

// type Model struct {
// 	ID         int `gorm:"primary_key" json:"id"`
// 	CreatedOn  int `json:"created_on"`
// 	ModifiedOn int `json:"modified_on"`
// 	DeletedOn  int `json:"deleted_on"`
// }

// Setup initializes the database instance
func Setup() {
	dsn := "root:yl123456@tcp(localhost:3306)/xitulu?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Fatalln("err", err)
	}
}
