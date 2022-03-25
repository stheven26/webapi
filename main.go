package main

import (
	"project/controller"
	"project/migrations"

	// "project/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//Config
	db := migrations.SetUpDatabase()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	r.POST("/api/v1/mahasiswa", controller.CreateDataMahasiswa)
	//GET All Data
	r.GET("/api/v1/mahasiswa", controller.GetDataMahasiswa)
	//Update Data >> Update Data
	r.PUT("/api/v1/mahasiswa/:nim", controller.UpdateDataMahasiswa)
	//Delete Data >> Delete data
	r.DELETE("/api/v1/mahasiswa", controller.DeleteDataMahasiswa)

	//POST Data >> Create Data
	r.POST("/api/v1/matakuliah", controller.CreateDataMatkul)
	//GET All Data
	r.GET("/api/v1/matakuliah", controller.GetDataMatkul)
	//Update Data >> Update Data
	r.PUT("/api/v1/matakuliah/:kode", controller.UpdateDataMatkul)
	//Delete Data >> Delete data
	r.DELETE("/api/v1/matakuliah", controller.DeleteDataMatkul)

	r.Run()
}
