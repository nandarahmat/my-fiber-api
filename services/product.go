package services

import (
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nandarahmat/my-fiber-api/database"
	"github.com/nandarahmat/my-fiber-api/models"
	"gorm.io/gorm"
)

func GetProducts(c *fiber.Ctx) error {
	// Ambil query parameters
	namaProduk := c.Query("nama_produk")
	limit := c.QueryInt("limit", 10)
	page := c.QueryInt("page", 1)
	categoryID := c.QueryInt("category_id")
	tokoID := c.QueryInt("toko_id")
	minHarga := c.QueryInt("min_harga")
	maxHarga := c.QueryInt("max_harga")

	var products []models.Product
	query := database.DB.Preload("Images")

	// Filtering
	if namaProduk != "" {
		query = query.Where("nama_produk LIKE ?", "%"+namaProduk+"%")
	}
	if categoryID > 0 {
		query = query.Where("id_category = ?", categoryID)
	}
	if tokoID > 0 {
		query = query.Where("id_toko = ?", tokoID)
	}
	if minHarga > 0 {
		query = query.Where("harga_konsumen >= ?", minHarga)
	}
	if maxHarga > 0 {
		query = query.Where("harga_konsumen <= ?", maxHarga)
	}

	// Pagination
	offset := (page - 1) * limit
	query = query.Limit(limit).Offset(offset)

	err := query.Find(&products).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal mengambil produk"})
	}

	return c.JSON(fiber.Map{
		"message":  "Produk ditemukan",
		"products": products,
	})
}

func GetProductById(c *fiber.Ctx) error {
	id := c.Params("id")

	var product models.Product
	err := database.DB.Preload("Images").First(&product, id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Produk tidak ditemukan"})
	}

	return c.JSON(fiber.Map{
		"message": "Produk ditemukan",
		"product": product,
	})
}

func CreateProduct(c *fiber.Ctx) error {
	// Parse form data
	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid form data"})
	}

	namaProduk := form.Value["nama_produk"][0]
	categoryIDStr := form.Value["category_id"][0]
	hargaReseller := form.Value["harga_reseller"][0]
	hargaKonsumen := form.Value["harga_konsumen"][0]
	stok := form.Value["stok"][0]
	deskripsi := form.Value["deskripsi"][0]
	toko := form.Value["id_toko"][0]

	stokInt, err := strconv.Atoi(stok)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid stock value"})
	}

	tokoID, err := strconv.Atoi(toko)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid stock value"})
	}

	// Convert categoryID from string to uint
	categoryID, err := strconv.ParseUint(categoryIDStr, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid category ID"})
	}

	// Generate slug
	slug := strings.ToLower(strings.ReplaceAll(namaProduk, " ", "-"))

	// Simpan produk ke database
	product := models.Product{
		NamaProduk:    namaProduk,
		Slug:          slug,
		HargaReseller: hargaReseller,
		HargaKonsumen: hargaKonsumen,
		Stok:          stokInt,
		Deskripsi:     deskripsi,
		IDCategory:    uint(categoryID), // Convert to uint
		IDToko:        uint(tokoID),
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	if err := database.DB.Create(&product).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create product"})
	}

	// Simpan foto produk
	photos := form.File["photos"]
	var productImages []models.ProductImage

	for _, photo := range photos {
		filePath := fmt.Sprintf("public/uploads/produk_%s", photo.Filename)
		if err := saveFile(photo, filePath); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to upload image"})
		}

		productImages = append(productImages, models.ProductImage{
			IDProduct: product.ID,
			URL:       "/" + filePath,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		})
	}

	// Insert images into database
	if len(productImages) > 0 {
		if err := database.DB.Create(&productImages).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save images"})
		}
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Product created successfully", "product": product})
}

func UpdateProduct(c *fiber.Ctx) error {
	// Get the product ID from the URL parameters
	productID := c.Params("id")

	// Parse form data
	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid form data"})
	}

	// Retrieve the existing product from the database
	var product models.Product
	if err := database.DB.First(&product, productID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Product not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve product"})
	}

	// Update product fields
	if len(form.Value["nama_produk"]) > 0 {
		namaProduk := form.Value["nama_produk"][0]
		product.NamaProduk = namaProduk
		product.Slug = strings.ToLower(strings.ReplaceAll(namaProduk, " ", "-"))
	}

	if len(form.Value["category_id"]) > 0 {
		categoryIDStr := form.Value["category_id"][0]
		categoryID, err := strconv.ParseUint(categoryIDStr, 10, 32)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid category ID"})
		}
		product.IDCategory = uint(categoryID)
	}

	if len(form.Value["harga_reseller"]) > 0 {
		product.HargaReseller = form.Value["harga_reseller"][0]
	}

	if len(form.Value["harga_konsumen"]) > 0 {
		product.HargaKonsumen = form.Value["harga_konsumen"][0]
	}

	stokInt, err := strconv.Atoi(form.Value["stok"][0])
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid stock value"})
	}

	if len(form.Value["stok"]) > 0 {
		product.Stok = stokInt
	}

	if len(form.Value["deskripsi"]) > 0 {
		product.Deskripsi = form.Value["deskripsi"][0]
	}

	product.UpdatedAt = time.Now()

	// Save the updated product to the database
	if err := database.DB.Save(&product).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update product"})
	}

	// Handle photo updates
	photos := form.File["photos"]
	var productImages []models.ProductImage

	for _, photo := range photos {
		filePath := fmt.Sprintf("public/uploads/produk_%s", photo.Filename)
		if err := saveFile(photo, filePath); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to upload image"})
		}

		productImages = append(productImages, models.ProductImage{
			IDProduct: product.ID,
			URL:       "/" + filePath,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		})
	}

	// Insert or update images in the database
	if len(productImages) > 0 {
		// Optionally, you can delete existing images or update them based on your requirements
		if err := database.DB.Where("id_product = ?", product.ID).Delete(&models.ProductImage{}).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete old images"})
		}

		if err := database.DB.Create(&productImages).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save images"})
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Product updated successfully", "product": product})
}

func DeleteProduct(c *fiber.Ctx) error {
	// Get the product ID from the URL parameters
	productID := c.Params("id")

	// Retrieve the existing product from the database
	var product models.Product
	if err := database.DB.First(&product, productID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Product not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve product"})
	}

	// Retrieve associated images
	var productImages []models.ProductImage
	if err := database.DB.Where("id_produk = ?", product.ID).Find(&productImages).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve product images"})
	}

	// Delete images from the filesystem
	for _, image := range productImages {
		imagePath := filepath.Join("public", "uploads", filepath.Base(image.URL)) // Ensure the path is correct

		// Check if the file exists
		if _, err := os.Stat(imagePath); os.IsNotExist(err) {
			continue // Skip to the next image if the file does not exist
		}

		// Attempt to delete the file
		if err := os.Remove(imagePath); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete image from filesystem"})
		}
	}

	// Delete the product images from the database
	if err := database.DB.Where("id_produk = ?", product.ID).Delete(&models.ProductImage{}).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete product images from database"})
	}

	// Delete the product from the database
	if err := database.DB.Delete(&product).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete product"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Product deleted successfully"})
}

func saveFile(file *multipart.FileHeader, filePath string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Create directory if not exists
	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	dst, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = dst.ReadFrom(src)
	return err
}
