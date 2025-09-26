package controllers

import (
	"siduk/models"
	"siduk/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type LoginRequest struct {
	NoWhatsapp string `json:"no_whatsapp"`
	PIN        string `json:"pin"`
}

func Login(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req LoginRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
		}
		var user models.Penduduk
		if err := db.Where("no_whatsapp = ?", req.NoWhatsapp).First(&user).Error; err != nil {
			return fiber.ErrUnauthorized
		}
		if !utils.CheckPIN(req.PIN, user.PIN) {
			return fiber.ErrUnauthorized
		}
		// generate JWT, set cookie/session
		token, err := utils.GenerateJWT(user.ID, user.Role)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Gagal generate token"})
		}
		return c.JSON(fiber.Map{"token": token})
	}
}

type SignupRequest struct {
	NoWhatsapp string `json:"no_whatsapp"`
	Nama       string `json:"nama"`
}

func Signup(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req SignupRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
		}
		// Generate random PIN, hash it, send PIN ke WhatsApp (dummy)
		pin := utils.GeneratePIN()
		utils.SendPINWhatsApp(req.NoWhatsapp, pin) // assume success
		hashedPIN, _ := utils.HashPIN(pin)
		user := models.Penduduk{
			Nama:       req.Nama,
			NoWhatsapp: req.NoWhatsapp,
			PIN:        hashedPIN,
			CreatedAt:  time.Now().Unix(),
			Role:       "user",
		}
		if err := db.Create(&user).Error; err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Gagal signup"})
		}
		return c.JSON(fiber.Map{"message": "PIN dikirim ke WhatsApp"})
	}
}
