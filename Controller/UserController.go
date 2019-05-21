package Controller

import (
	"github.com/jinzhu/gorm"
	"github.com/yamagn/gallari-back/Model"
)

func Create(db *gorm.DB, user Model.User) {
	db.Create(&user)
}

func GetOne(db *gorm.DB, id int) Model.User{
	findedUser := Model.User{}
	findedUser.Id = id
	db.First(&findedUser)
	return findedUser
}

func GetAll(db *gorm.DB) []Model.User {
	users := []Model.User{}
	db.Find(&users)
	return users
}

func Delete(db *gorm.DB, id int) {
	user := Model.User{}
	user.Id = id
	db.First(&user)
	db.Delete(&user)
}

func Update(db *gorm.DB, user Model.User) {
	userBefore := Model.User{}
	userBefore.Id = user.Id
	db.Model(userBefore).Update(user)
}
