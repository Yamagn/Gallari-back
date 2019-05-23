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

func GetOne(db *gorm.DB, id int) Model.ReturnUser{
	user := Model.User{}
	db.Where("ID = ?", id).First(&user)
	returnUser := Model.ReturnUser{}
	returnUser.Id = user.Id
	returnUser.DisplayName = user.DisplayName
	returnUser.Email = user.Email
	returnUser.CreatedAt = user.CreatedAt
	returnUser.UpdatedAt = user.UpdatedAt
	return returnUser
}

func GetAll(db *gorm.DB) []Model.ReturnUser {
	var users []Model.User
	db.Find(&users)
	var returnUsers []Model.ReturnUser
	for _, user := range users{
		returnUser := Model.ReturnUser{}
		returnUser.Id = user.Id
		returnUser.DisplayName = user.DisplayName
		returnUser.Email = user.Email
		returnUser.CreatedAt = user.CreatedAt
		returnUser.UpdatedAt = user.UpdatedAt
		returnUsers = append(returnUsers, returnUser)
	}

	return returnUsers
}

func Delete(db *gorm.DB, id int) {
	user := Model.User{}
	db.Where("ID = ?", id).Delete(&user)
}

func Update(db *gorm.DB, id int, data Model.User) {
	user := Model.User{}
	db.Where("ID = ?", id).First(&user).Update(&data)
}
