package repository

import (
	"github.com/inventario-go/alejoab12/user-api/api/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func InitialDB() {
	db := openConnection()
	defer db.Close()
	db.AutoMigrate(&models.User{}, &models.DocumentType{})
}
func FindUserById(id string) models.User {
	db := openConnection()
	defer db.Close()
	var user models.User
	db.Find(&user, db.Where("id=?", id))
	return user
}
func CreateUser(user *models.User) {
	db := openConnection()
	defer db.Close()
	db.Create(*user)
}

func openConnection() *gorm.DB {
	db, error := gorm.Open("mysql", "root@tcp(127.0.0.1:3306)/user")
	if error != nil {
		panic("Failed connect to DB")
	}
	db.SingularTable(true)
	return db
}
