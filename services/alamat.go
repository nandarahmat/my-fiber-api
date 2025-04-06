package services

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nandarahmat/my-fiber-api/database"
	"github.com/nandarahmat/my-fiber-api/models"
)

func GetUserAlamat(c *fiber.Ctx) error {
	// Ambil userID dari token
	userIDFloat, ok := c.Locals("userID").(float64)
	if !ok {
		return c.Status(400).JSON(fiber.Map{"error": "Gagal mendapatkan userID dari token"})
	}
	userID := uint(userIDFloat)

	// Ambil semua alamat berdasarkan userID
	var alamat []models.Alamat
	if err := database.DB.Where("id_user = ?", userID).Find(&alamat).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal mengambil data alamat"})
	}

	return c.JSON(alamat)
}

func GetUserAlamatById(c *fiber.Ctx) error {
	// Ambil userID dari token
	userIDFloat, ok := c.Locals("userID").(float64)
	if !ok {
		return c.Status(400).JSON(fiber.Map{"error": "Gagal mendapatkan userID dari token"})
	}
	userID := uint(userIDFloat)

	// Ambil id alamat dari parameter
	alamatID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "ID alamat tidak valid"})
	}

	// Cari alamat berdasarkan id dan user_id
	var alamat models.Alamat
	if err := database.DB.Where("id = ? AND id_user = ?", alamatID, userID).First(&alamat).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Alamat tidak ditemukan atau bukan milik Anda"})
	}

	return c.JSON(alamat)
}

func CreateUserAlamat(c *fiber.Ctx) error {
	// Ambil userID dari token
	userIDFloat, ok := c.Locals("userID").(float64)
	if !ok {
		return c.Status(400).JSON(fiber.Map{"error": "Gagal mendapatkan userID dari token"})
	}
	userID := uint(userIDFloat)

	// Struktur untuk menangkap data request
	var alamat models.Alamat
	if err := c.BodyParser(&alamat); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Gagal parsing data"})
	}

	// Set id_user dari token
	alamat.IDUser = userID
	alamat.CreatedAt = time.Now()
	alamat.UpdatedAt = time.Now()

	// Simpan ke database
	if err := database.DB.Create(&alamat).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal menyimpan alamat"})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "Alamat berhasil ditambahkan",
		"alamat":  alamat,
	})
}

func UpdateUserAlamat(c *fiber.Ctx) error {
	// Ambil userID dari token
	userIDFloat, ok := c.Locals("userID").(float64)
	if !ok {
		return c.Status(400).JSON(fiber.Map{"error": "Gagal mendapatkan userID dari token"})
	}
	userID := uint(userIDFloat)

	// Ambil ID alamat dari parameter
	alamatID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "ID alamat tidak valid"})
	}

	// Cari alamat berdasarkan id dan id_user
	var alamat models.Alamat
	if err := database.DB.Where("id = ? AND id_user = ?", alamatID, userID).First(&alamat).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Alamat tidak ditemukan"})
	}

	// Parsing data request
	if err := c.BodyParser(&alamat); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Gagal parsing data"})
	}

	// Update waktu perubahan
	alamat.UpdatedAt = time.Now()

	// Simpan perubahan
	if err := database.DB.Save(&alamat).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal memperbarui alamat"})
	}

	return c.JSON(fiber.Map{
		"message": "Alamat berhasil diperbarui",
		"alamat":  alamat,
	})
}

func DeleteUserAlamat(c *fiber.Ctx) error {
	// Ambil userID dari token
	userIDFloat, ok := c.Locals("userID").(float64)
	if !ok {
		return c.Status(400).JSON(fiber.Map{"error": "Gagal mendapatkan userID dari token"})
	}
	userID := uint(userIDFloat)

	// Ambil ID alamat dari parameter
	alamatID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "ID alamat tidak valid"})
	}

	// Cari alamat berdasarkan id dan id_user
	var alamat models.Alamat
	if err := database.DB.Where("id = ? AND id_user = ?", alamatID, userID).First(&alamat).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Alamat tidak ditemukan atau bukan milik Anda"})
	}

	// Hapus alamat
	if err := database.DB.Delete(&alamat).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal menghapus alamat"})
	}

	return c.JSON(fiber.Map{
		"message": "Alamat berhasil dihapus",
	})
}
