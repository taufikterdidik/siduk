package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"siduk/models"
)

func StatistikPublik(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var total, laki, perempuan int64
		db.Model(&models.Penduduk{}).Count(&total)
		db.Model(&models.Penduduk{}).Where("jenis_kelamin = ?", "Laki-laki").Count(&laki)
		db.Model(&models.Penduduk{}).Where("jenis_kelamin = ?", "Perempuan").Count(&perempuan)
		// Data dummy usia: hitung per range umur bisa ditambah
		return c.Render("statistik", fiber.Map{
			"Total":     total,
			"Laki":      laki,
			"Perempuan": perempuan,
		})
	}
}