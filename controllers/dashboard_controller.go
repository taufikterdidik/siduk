package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Dashboard (AdminLTE): hanya untuk admin/operator yang sudah login
func Dashboard(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userRole := c.Locals("role")
		if userRole != "admin" && userRole != "operator" {
			return c.Status(403).SendString("Forbidden")
		}
		// Data statistik ringkas untuk dashboard, bisa dikembangkan
		var total int64
		db.Model(&struct{ID uint}{}).Table("penduduks").Count(&total)
		return c.Render("dashboard", fiber.Map{
			"TotalPenduduk": total,
			"Role":          userRole,
		})
	}
}