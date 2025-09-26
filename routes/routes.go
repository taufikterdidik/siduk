package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"siduk/controllers"
	"siduk/middleware"
)

// SetupRoutes sets up all application routes
func SetupRoutes(app *fiber.App, db *gorm.DB) {
	// Public routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/statistik")
	})
	app.Get("/statistik", controllers.StatistikPublik(db))
	app.Static("/public", "./public") // serve static files for AdminLTE, JS, CSS

	// API v1 group
	api := app.Group("/api")

	// Auth routes
	api.Post("/login", controllers.Login(db))
	api.Post("/signup", controllers.Signup(db))
	api.Post("/reset-pin", controllers.ResetPIN(db)) // opsional: implementasi reset PIN

	// Statistik API (untuk frontend JS chart, dsb)
	api.Get("/statistik/summary", controllers.StatistikSummary(db)) // data JSON summary
	api.Get("/statistik/usia", controllers.StatistikUsia(db))
	api.Get("/statistik/jk", controllers.StatistikJK(db))

	// Protected routes (JWT middleware, hanya untuk user login)
	admin := api.Group("/admin", middleware.JWTProtected()) // implementasi JWTProtected di middleware/jwt_middleware.go
	// CRUD Penduduk
	admin.Get("/penduduk", controllers.ListPenduduk(db))
	admin.Get("/penduduk/:id", controllers.GetPenduduk(db))
	admin.Post("/penduduk", controllers.CreatePenduduk(db))
	admin.Put("/penduduk/:id", controllers.UpdatePenduduk(db))
	admin.Delete("/penduduk/:id", controllers.DeletePenduduk(db))
	// Upload Excel/CSV (opsional)
	admin.Post("/penduduk/import", controllers.ImportPendudukCSV(db))
	// CRUD Keluarga
	admin.Get("/keluarga", controllers.ListKeluarga(db))
	admin.Get("/keluarga/:id", controllers.GetKeluarga(db))
	admin.Post("/keluarga", controllers.CreateKeluarga(db))
	admin.Put("/keluarga/:id", controllers.UpdateKeluarga(db))
	admin.Delete("/keluarga/:id", controllers.DeleteKeluarga(db))

	// AdminLTE dashboard (protected)
	app.Get("/dashboard", middleware.JWTWebProtected(), controllers.Dashboard(db))
	app.Get("/penduduk", middleware.JWTWebProtected(), controllers.PendudukPage(db))
	app.Get("/keluarga", middleware.JWTWebProtected(), controllers.KeluargaPage(db))
}