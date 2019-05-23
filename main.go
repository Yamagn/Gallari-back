package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/yamagn/gallari-back/Conf"
	"github.com/yamagn/gallari-back/Model"
	"github.com/yamagn/gallari-back/Router"
)

func main() {
	db := Conf.GormConnect()
	db.Set("gorm:table_options", "ENGINE=InnoDB")
	db.AutoMigrate(&Model.User{})
	db.AutoMigrate(&Model.Product{})
	db.LogMode(true)
	defer db.Close()

	r := Router.Router(db)
	r.Run()
}