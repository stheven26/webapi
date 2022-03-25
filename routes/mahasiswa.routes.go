package routes

import (
	"project/controller"

	"github.com/gin-gonic/gin"
)

func MahasiswaRoutes() {

	r := gin.Default()
	//POST Data >> Create Data
	r.POST("/api/v1/mahasiswa", controller.CreateDataMahasiswa)
	//GET All Data
	r.GET("/api/v1/mahasiswa", controller.GetDataMahasiswa)
	//Update Data >> Update Data
	r.PUT("/api/v1/mahasiswa/:nim", controller.UpdateDataMahasiswa)
	//Delete Data >> Delete data
	r.DELETE("/api/v1/mahasiswa", controller.DeleteDataMahasiswa)
}
