package Router

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/yamagn/gallari-back/Controller/ProductController"
	"github.com/yamagn/gallari-back/Controller/UserController"
	"github.com/yamagn/gallari-back/Model"
	"net/http"
	"strconv"
	"time"
)

func Router(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	// product
	r.POST("/product/create", func(c *gin.Context) {
		var req Model.Product
		err := c.BindJSON(&req)
		req.CreatedAt = time.Now()
		req.UpdatedUt = time.Now()
		req.Id = 0
		if err != nil {
			c.String(http.StatusBadRequest, "Request is failed: "+err.Error())
		}
		ProductController.Create(db, req)
		if db.NewRecord(req) == false {
			c.JSON(http.StatusOK, req)
		}
	})
	r.GET("/product/find/:Id", func(c *gin.Context) {
		param := c.Param("Id")
		id, e := strconv.Atoi(param)
		if e != nil {
			c.String(http.StatusBadRequest, "Request is failed: " + e.Error())
		}
		product := ProductController.GetOne(db, id)
		jsonBytes, err := json.Marshal(product)
		if err != nil {
			c.String(http.StatusBadRequest, "Request is failed: "+ err.Error())
			return
		}
		c.JSON(http.StatusOK, jsonBytes)
	})
	r.GET("/product/all", func(c *gin.Context) {
		products := ProductController.GetAll(db)
		jsonBytes, err := json.Marshal(products)
		if err != nil {
			c.String(http.StatusInternalServerError, "Error: " + err.Error())
			return
		}
		c.JSON(http.StatusOK, jsonBytes)
	})
	r.DELETE("/product/delete/:Id", func(c *gin.Context) {
		param := c.Param("Id")
		id, e := strconv.Atoi(param)
		if e != nil {
			c.String(http.StatusBadRequest, "Request is failed: " + e.Error())
		}
		ProductController.Delete(db, id)
	})
	r.PUT("/product/update", func(c *gin.Context) {
		c.String(200, "Hello World!")
	})

	// user
	r.POST("/user/create", func(c *gin.Context) {
		var user Model.User
		err := c.BindJSON(&user)
		if err != nil {
			c.String(http.StatusBadRequest, "Request is failed: "+err.Error())
		}
		user.CreatedAt = time.Now()
		user.UpdatedAt = time.Now()
		user.Id = 0
		UserController.Create(db, user)
		if db.NewRecord(user) == false {
			c.JSON(http.StatusOK, user)
		}
	})
	r.GET("/user/find/:Id", func(c *gin.Context) {
		param := c.Param("Id")
		id, e := strconv.Atoi(param)
		if e != nil {
			c.String(http.StatusBadRequest, "Request is failed: " + e.Error())
		}
		user := UserController.GetOne(db, id)
		c.JSON(http.StatusOK, user)
	})
	r.GET("/user/all", func(c *gin.Context) {
		users := UserController.GetAll(db)
		c.JSON(http.StatusOK, users)
	})
	r.DELETE("/user/delete/:Id", func(c *gin.Context) {
		param := c.Param("Id")
		id, e := strconv.Atoi(param)
		if e != nil {
			c.String(http.StatusBadRequest, "Request is failed: " + e.Error())
		}
		UserController.Delete(db, id)
	})
	r.PUT("/user/update", func(c *gin.Context) {
		c.String(200, "Hello World!")
	})

	return r
}
