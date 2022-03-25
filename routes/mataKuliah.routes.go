package routes

import (
	"project/controller"

	"github.com/gin-gonic/gin"
)

func MataKuliahRoutes() {

	r := gin.Default()
	//POST Data >> Create Data
	r.POST("/api/v1/matakuliah", controller.CreateDataMatkul)
	//GET All Data
	r.GET("/api/v1/matakuliah", controller.GetDataMatkul)
	//Update Data >> Update Data
	r.PUT("/api/v1/matakuliah/:kode", controller.UpdateDataMatkul)
	//Delete Data >> Delete data
	r.DELETE("/api/v1/matakuliah", controller.DeleteDataMatkul)
}
