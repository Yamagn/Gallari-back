package ProductController

import (
	"github.com/jinzhu/gorm"
	"github.com/yamagn/gallari-back/Model"
)

func Create(db *gorm.DB, product Model.Product) {
	db.NewRecord(product)
	db.Create(&product)
}

func GetOne(db *gorm.DB, id int) Model.Product{
	foundProduct := Model.Product{}
	db.Where("Id = ?", id).First(&foundProduct)
	return foundProduct
}

func GetAll(db *gorm.DB) []Model.Product {
	var products []Model.Product
	db.Find(&products)
	return products
}

func Delete(db *gorm.DB, id int) {
	product := Model.Product{}
	db.Where("Id = ?", id).First(&product)
	db.Delete(&product)
}

func Update(db *gorm.DB, id int, data Model.Product) {
	user := Model.User{}
	db.Where("ID = ?", id).First(&user).Update(&data)
}