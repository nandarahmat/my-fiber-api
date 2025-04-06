package services

import (
	"errors"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nandarahmat/my-fiber-api/database"
	"github.com/nandarahmat/my-fiber-api/models"
	"gorm.io/gorm"
)

func GetAllTrx(c *fiber.Ctx) error {
	// Ambil UserID dari token
	userIDFloat, ok := c.Locals("userID").(float64)
	if !ok {
		return c.Status(400).JSON(fiber.Map{"error": "Gagal mendapatkan userID dari token"})
	}
	userID := uint(userIDFloat)

	// Ambil query parameter untuk pagination
	page := c.QueryInt("page", 1)
	pageSize := c.QueryInt("pageSize", 10)
	offset := (page - 1) * pageSize

	var trxList []models.Trx
	query := database.DB.Model(&models.Trx{}).Where("id_user = ?", userID)

	if err := query.Offset(offset).Limit(pageSize).Find(&trxList).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(trxList)
}

func GetTrxByID(c *fiber.Ctx) error {
	trxID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "ID tidak valid"})
	}

	var trx models.Trx
	if err := database.DB.First(&trx, trxID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(404).JSON(fiber.Map{"error": "Transaksi tidak ditemukan"})
		}
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(trx)
}

// StoreTrx menambahkan transaksi baru
func StoreTrx(c *fiber.Ctx) error {
	type DetailInput struct {
		IDProduk  uint `json:"id_produk"`
		Kuantitas int  `json:"kuantitas"`
	}

	type TrxInput struct {
		AlamatPengiriman uint          `json:"alamat_pengiriman"`
		MethodBayar      string        `json:"method_bayar"`
		Details          []DetailInput `json:"details"`
	}

	var input TrxInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	userIDFloat, ok := c.Locals("userID").(float64)
	if !ok {
		return c.Status(400).JSON(fiber.Map{"error": "Gagal mendapatkan userID dari token"})
	}
	userID := uint(userIDFloat)

	var hargaTotal int
	var details []models.DetailTrx

	for _, detail := range input.Details {
		var product models.Product
		if err := database.DB.First(&product, detail.IDProduk).Error; err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Produk tidak ditemukan"})
		}

		hargaReseller, err := strconv.Atoi(product.HargaReseller)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Harga reseller tidak valid"})
		}

		hargaKonsumen, err := strconv.Atoi(product.HargaKonsumen)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Harga konsumen tidak valid"})
		}

		// Buat log produk
		logProduk := models.LogProduk{
			IDProduk:      product.ID,
			NamaProduk:    product.NamaProduk,
			Slug:          product.Slug,
			HargaReseller: hargaReseller,
			HargaKonsumen: hargaKonsumen,
			Deskripsi:     product.Deskripsi,
			IDToko:        product.IDToko,
			IDCategory:    product.IDCategory,
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		}

		if err := database.DB.Create(&logProduk).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Gagal menyimpan log produk"})
		}

		subtotal := hargaKonsumen * detail.Kuantitas
		hargaTotal += subtotal

		details = append(details, models.DetailTrx{
			IDLogProduk: logProduk.ID,
			IDToko:      logProduk.IDToko,
			Kuantitas:   detail.Kuantitas,
			HargaTotal:  subtotal,
		})
	}

	trx := models.Trx{
		UserID:           userID,
		AlamatPengiriman: input.AlamatPengiriman,
		HargaTotal:       hargaTotal,
		KodeInvoice:      generateInvoiceCode(),
		MethodBayar:      input.MethodBayar,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	if err := database.DB.Create(&trx).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal menyimpan transaksi"})
	}

	for i := range details {
		details[i].IDTrx = trx.ID
	}
	if err := database.DB.Create(&details).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal menyimpan detail transaksi"})
	}

	return c.JSON(trx)
}

func generateInvoiceCode() string {
	return "INV-" + time.Now().Format("20060102150405")
}
