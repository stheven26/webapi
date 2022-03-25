package controller

import (
	"net/http"
	"project/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type MahasiswaInput struct {
	ID       int    `json:"id" binding:"required"`
	Nama     string `json:"nama" binding:"required"`
	Prodi    string `json:"prodi" binding:"required"`
	Fakultas string `json:"fakultas" binding:"required"`
	NIM      int    `json:"nim" binding:"required"`
}

func CreateDataMahasiswa(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	//validasi inputan
	var dataInput MahasiswaInput
	if err := c.ShouldBindJSON(&dataInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	//	prose input data
	mhs := models.Mahasiswa{
		ID:       dataInput.ID,
		Nama:     dataInput.Nama,
		Prodi:    dataInput.Prodi,
		Fakultas: dataInput.Fakultas,
		NIM:      dataInput.NIM,
	}

	db.Create(&mhs)

	//	menampilkan hasil
	c.JSON(http.StatusOK, gin.H{
		"Message": "Sukses input data Mahasiswa",
		"Data":    mhs,
		"time":    time.Now(),
	})
}

//GET Data
func GetDataMahasiswa(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var mhs []models.Mahasiswa
	db.Find(&mhs)
	c.JSON(http.StatusOK, gin.H{
		"data": mhs,
		"time": time.Now(),
	})
}

//UPDATE Data >> Update Data
func UpdateDataMahasiswa(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//cek data
	var mhs models.Mahasiswa
	if err := db.Where("nim = ?", c.Param("nim")).First(&mhs).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Data mahasiswa tidak di temukan",
		})
		return
	}

	//validasi inputan
	var dataInput MahasiswaInput
	if err := c.ShouldBindJSON(&dataInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	//	prose Ubah data
	db.Model(&mhs).Update(&dataInput)

	//	menampilkan hasil
	c.JSON(http.StatusOK, gin.H{
		"Message": "Sukses mengubah data Mahasiswa",
		"Data":    mhs,
		"time":    time.Now(),
	})
}

// Delete Data >> Hapus Data
func DeleteDataMahasiswa(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//cek data
	var mhs models.Mahasiswa
	if err := db.Where("nim = ?", c.Query("nim")).First(&mhs).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Data mahasiswa tidak di temukan",
		})
		return
	}
	//	prose hapus data
	db.Delete(&mhs)

	//	menampilkan hasil
	c.JSON(http.StatusOK, gin.H{
		"Data": true,
	})
}
