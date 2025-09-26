package controllers

import (
	"siduk/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// ListPenduduk: GET /api/admin/penduduk
func ListPenduduk(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var penduduks []models.Penduduk
		db.Preload("Keluarga").Find(&penduduks)
		return c.JSON(penduduks)
	}
}

// CRUD lain: GetPenduduk, CreatePenduduk, UpdatePenduduk, DeletePenduduk