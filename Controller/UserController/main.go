package UserController

import (
	"github.com/jinzhu/gorm"
	"github.com/yamagn/gallari-back/Model"
	"golang.org/x/crypto/bcrypt"
)

func passwordHash(pw string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	if err != nil {
		return "", nil
	}
	return string(hash), err
}

func Create(db *gorm.DB, user Model.User) {
	hash, err := passwordHash(user.Password)
	if err != nil {
		panic(err)
	}
	user.Password = hash
	db.NewRecord(user)
	db.Create(&user)
}

func GetOne(db *gorm.DB, id int) Model.User{
	user := Model.User{}
	db.Where("ID = ?", id).First(&user)
	return user
}

func GetAll(db *gorm.DB) []Model.User {
	var users []Model.User
	db.Find(&users)
	return users
}

func Delete(db *gorm.DB, id int) {
	user := Model.User{}
	db.Where("ID = ?", id).Delete(&user)
}

func Update(db *gorm.DB, id int, data Model.User) {
	user := Model.User{}
	db.Where("ID = ?", id).First(&user).Update(&data)
}
