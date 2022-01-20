package dao

import (
	"Project01/main/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserDao interface {
	Insert(user models.User)
	Update(user models.User)
	Delete(user models.User)
	FindAll() []models.User
	CloseDB()
}

type Database struct {
	connection *gorm.DB
}

func (db *Database) Insert(user models.User) {
	db.connection.Create(&user)
}

func (db *Database) Update(user models.User) {
	db.connection.Save(&user)
}

func (db *Database) Delete(user models.User) {
	db.connection.Delete(&user)
}

func (db *Database) FindAll() []models.User {
	var userList []models.User
	db.connection.Find(&userList)
	return userList
}

func InitUserDao() UserDao {
	dsn := "root:software@tcp(127.0.0.1:3306)/amazon?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		return nil
	}
	return &Database{connection: db}
}

func (db *Database) CloseDB() {
	database, _ := db.connection.DB()
	err := database.Close()
	if err != nil {
		panic("Failed to close database")
	}
}
