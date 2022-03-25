package migrations

import (
	"project/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func SetUpDatabase() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@(localhost)/webgolangdb?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("Error koneksi kedalam database")
	}

	db.AutoMigrate(&models.Mahasiswa{})
	db.AutoMigrate(&models.MataKuliah{})

	return db
}
