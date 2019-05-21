package Controller

import (
	"github.com/jinzhu/gorm"
	"github.com/yamagn/gallari-back/Model"
)

func Create(db *gorm.DB, product Model.Product) {
	db.Create(&product)

}

func GetOne(db *gorm.DB, id int) Model.Product{
	findedProduct := Model.Product{}
	findedProduct.ProductId = id
	db.First(&findedProduct)
	return findedProduct
}

func GetAll(db *gorm.DB) []Model.Product {
	products := []Model.Product{}
	db.Find(&products)
	return products
}

func Delete(db *gorm.DB, id int) {
	product := Model.Product{}
	product.ProductId = id
	db.First(&product)
	db.Delete(&product)
}

func Update(db *gorm.DB, product Model.Product) {
	productBefore := Model.Product{}
	productBefore.ProductId = product.ProductId
	db.Model(productBefore).Update(product)
}