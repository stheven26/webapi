package controller

import (
	"net/http"
	"project/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type MataKuliahInput struct {
	ID    int    `json:"id" binding:"required"`
	Kode  string `json:"kode" binding:"required"`
	Nama  string `json:"nama" binding:"required"`
	SKS   int    `json:"sks" binding:"required"`
	Dosen string `json:"dosen" binding:"required"`
}

func CreateDataMatkul(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	//validasi inputan
	var dataInput MataKuliahInput
	if err := c.ShouldBindJSON(&dataInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	//	prose input data
	matkul := models.MataKuliah{
		ID:    dataInput.ID,
		Kode:  dataInput.Kode,
		Nama:  dataInput.Nama,
		SKS:   dataInput.SKS,
		Dosen: dataInput.Dosen,
	}

	db.Create(&matkul)

	//	menampilkan hasil
	c.JSON(http.StatusOK, gin.H{
		"Message": "Sukses input data Mata Kuliah",
		"Data":    matkul,
		"time":    time.Now(),
	})
}

//GET Data
func GetDataMatkul(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var matkul []models.MataKuliah
	db.Find(&matkul)
	c.JSON(http.StatusOK, gin.H{
		"data": matkul,
		"time": time.Now(),
	})
}

//UPDATE Data >> Update Data
func UpdateDataMatkul(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//cek data
	var matkul models.MataKuliah
	if err := db.Where("kode = ?", c.Param("kode")).First(&matkul).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Data Mata Kuliah tidak di temukan",
		})
		return
	}

	//validasi inputan
	var dataInput MataKuliahInput
	if err := c.ShouldBindJSON(&dataInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	//	prose Ubah data
	db.Model(&matkul).Update(&dataInput)

	//	menampilkan hasil
	c.JSON(http.StatusOK, gin.H{
		"Message": "Sukses mengubah data Mata Kuliah",
		"Data":    matkul,
		"time":    time.Now(),
	})
}

// Delete Data >> Hapus Data
func DeleteDataMatkul(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//cek data
	var matkul models.MataKuliah
	if err := db.Where("kode = ?", c.Query("kode")).First(&matkul).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Data Mata Kuliah tidak di temukan",
		})
		return
	}
	//	prose hapus data
	db.Delete(&matkul)

	//	menampilkan hasil
	c.JSON(http.StatusOK, gin.H{
		"Data": true,
	})
}
