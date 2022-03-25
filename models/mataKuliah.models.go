package models

type MataKuliah struct {
	ID    int    `json:"id"`
	Kode  string `json:"kode" gorm:"primary_key"`
	Nama  string `json:"nama"`
	SKS   int    `json:"sks"`
	Dosen string `json:"dosen"`
}
