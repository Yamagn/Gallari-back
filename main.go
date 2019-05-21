package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yamagn/gallari-back/Conf"
	"github.com/yamagn/gallari-back/Model"
	"github.com/yamagn/gallari-back/Controller"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db := Conf.GormConnect()
	db.Set("gorm:table_options", "ENGINE=mysql")
	db.AutoMigrate(&Model.User{})
	db.AutoMigrate(&Model.Product{})
	defer db.Close()

	r := gin.Default()

	// product
	r.POST("/product/create", func(c *gin.Context) {
		c.String(200, "Hello World!")
	})
	r.GET("/product/:Id", func(c *gin.Context) {
		c.String(200, "Hello World!")
	})
	r.GET("/product/all", func(c *gin.Context) {
		c.String(200, "Hello World!")
	})
	r.DELETE("/product/:Id", func(c *gin.Context) {
		c.String(200, "Hello World!")
	})
	r.PUT("/product/update", func(c *gin.Context) {
		c.String(200, "Hello World!")
	})

	// user
	r.POST("/user/create", func(c *gin.Context) {
		c.String(200, "Hello World!")
	})
	r.GET("/user/:Id", func(c *gin.Context) {
		c.String(200, "Hello World!")
	})
	r.GET("/user/all", func(c *gin.Context) {
		c.String(200, "Hello World!")
	})
	r.DELETE("/user/:Id", func(c *gin.Context) {
		c.String(200, "Hello World!")
	})
	r.PUT("/user/update", func(c *gin.Context) {
		c.String(200, "Hello World!")
	})

	r.Run()
}