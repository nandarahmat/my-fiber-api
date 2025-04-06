package services

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nandarahmat/my-fiber-api/database"
	"github.com/nandarahmat/my-fiber-api/models"
)

func GetMyToko(c *fiber.Ctx) error {
	// Ambil userID dari token
	userIDFloat, ok := c.Locals("userID").(float64)
	if !ok {
		return c.Status(400).JSON(fiber.Map{"error": "Gagal mendapatkan userID dari token"})
	}
	userID := uint(userIDFloat)

	// Ambil data toko berdasarkan id_user
	var toko models.Toko
	if err := database.DB.Where("id_user = ?", userID).First(&toko).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Toko tidak ditemukan"})
	}

	// Kirim response
	return c.JSON(toko)
}

func UpdateMyToko(c *fiber.Ctx) error {
	// Ambil userID dari token
	userIDFloat, ok := c.Locals("userID").(float64)
	if !ok {
		return c.Status(400).JSON(fiber.Map{"error": "Gagal mendapatkan userID dari token"})
	}
	userID := uint(userIDFloat)

	// Cek apakah toko milik user ada
	var toko models.Toko
	if err := database.DB.Where("id_user = ?", userID).First(&toko).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Toko tidak ditemukan"})
	}

	// Ambil input nama_toko dari form-data
	namaToko := c.FormValue("nama_toko")
	if namaToko != "" {
		toko.NamaToko = namaToko
	}

	// Cek apakah ada file gambar yang diupload
	file, err := c.FormFile("photo")
	if err == nil {
		// Buat folder "public/uploads" jika belum ada
		saveDir := "./public/uploads"
		if _, err := os.Stat(saveDir); os.IsNotExist(err) {
			os.MkdirAll(saveDir, os.ModePerm)
		}

		// Hapus gambar lama jika ada
		if toko.UrlFoto != "" {
			oldFilePath := "." + toko.UrlFoto // Path lama
			if _, err := os.Stat(oldFilePath); err == nil {
				os.Remove(oldFilePath) // Hapus file lama
			}
		}

		// Generate nama file unik berdasarkan timestamp
		ext := filepath.Ext(file.Filename)
		fileName := fmt.Sprintf("toko_%d_%d%s", userID, time.Now().Unix(), ext)
		filePath := filepath.Join(saveDir, fileName)

		// Simpan file ke folder public
		if err := c.SaveFile(file, filePath); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Gagal menyimpan gambar"})
		}

		// Buat URL akses gambar baru
		toko.UrlFoto = "/public/uploads/" + fileName
	}

	// Simpan perubahan ke database
	if err := database.DB.Save(&toko).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal memperbarui toko"})
	}

	// Kirim response
	return c.JSON(fiber.Map{
		"message": "Toko berhasil diperbarui",
		"toko":    toko,
	})
}

func GetAllToko(c *fiber.Ctx) error {
	// Ambil query parameter
	limit, _ := strconv.Atoi(c.Query("limit", "10")) // Default 10 data per halaman
	page, _ := strconv.Atoi(c.Query("page", "1"))    // Default halaman 1
	nama := c.Query("nama")                          // Filter berdasarkan nama toko

	// Hitung offset untuk pagination
	offset := (page - 1) * limit

	// Query ke database
	var tokoList []models.Toko
	query := database.DB.Limit(limit).Offset(offset)

	// Jika ada filter berdasarkan nama toko
	if nama != "" {
		query = query.Where("nama_toko LIKE ?", "%"+nama+"%")
	}

	// Eksekusi query
	if err := query.Find(&tokoList).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal mengambil data toko"})
	}

	// Hitung total data
	var total int64
	database.DB.Model(&models.Toko{}).Count(&total)

	// Kirim response
	return c.JSON(fiber.Map{
		"message": "Data toko berhasil diambil",
		"page":    page,
		"limit":   limit,
		"total":   total,
		"data":    tokoList,
	})
}

func GetTokoByID(c *fiber.Ctx) error {
	// Ambil parameter ID dari URL
	tokoID := c.Params("id")

	// Validasi apakah ID berupa angka
	if _, err := strconv.Atoi(tokoID); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "ID toko tidak valid"})
	}

	// Query ke database
	var toko models.Toko
	if err := database.DB.Where("id = ?", tokoID).First(&toko).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Toko tidak ditemukan"})
	}

	// Kirim response
	return c.JSON(fiber.Map{
		"message": "Data toko berhasil diambil",
		"data":    toko,
	})
}
