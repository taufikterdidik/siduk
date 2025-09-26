package models

import (
	"gorm.io/gorm"
)

type Penduduk struct {
	ID           uint   `gorm:"primaryKey"`
	Nama         string `gorm:"not null"`
	NIK          string `gorm:"uniqueIndex;size:16"`
	NoWhatsapp   string `gorm:"unique;not null"`
	PIN          string `json:"-"` // hashed
	JenisKelamin string
	TanggalLahir string // YYYY-MM-DD
	KeluargaID   uint
	Keluarga     Keluarga
	Role         string `gorm:"default:user"`
	CreatedAt    int64
	UpdatedAt    int64
}

func AutoMigratePenduduk(db *gorm.DB) error {
	return db.AutoMigrate(&Penduduk{})
}