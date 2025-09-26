package models

import (
	"gorm.io/gorm"
)

type Keluarga struct {
	ID        uint   `gorm:"primaryKey"`
	NoKK      string `gorm:"uniqueIndex"`
	Alamat    string
	Desa      string
	RT        string
	RW        string
	Penduduks []Penduduk
	CreatedAt int64
	UpdatedAt int64
}

func AutoMigrateKeluarga(db *gorm.DB) error {
	return db.AutoMigrate(&Keluarga{})
}