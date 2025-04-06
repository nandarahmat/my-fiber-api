package services

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"github.com/nandarahmat/my-fiber-api/database"
	"github.com/nandarahmat/my-fiber-api/models"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	var user models.User

	// Parsing body request
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Gagal parsing data"})
	}

	// Konversi tanggal lahir ke format MySQL (YYYY-MM-DD)
	parsedDate, err := time.Parse("02/01/2006", user.TanggalLahir)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Format tanggal harus DD/MM/YYYY"})
	}
	user.TanggalLahir = parsedDate.Format("2006-01-02")

	// Hash password sebelum disimpan
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.KataSandi), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal mengenkripsi password"})
	}
	user.KataSandi = string(hashedPassword)

	// Simpan user ke database
	if err := database.DB.Create(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal menyimpan data user"})
	}

	// **Tambahkan data toko setelah registrasi berhasil**
	toko := models.Toko{
		IDUser:   user.ID,
		NamaToko: user.Nama + " TOKO",
		UrlFoto:  "",
	}

	if err := database.DB.Create(&toko).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal menyimpan data toko"})
	}

	// Beri response sukses
	return c.Status(201).JSON(fiber.Map{
		"message": "Registrasi berhasil",
		"user": fiber.Map{
			"id":            user.ID,
			"nama":          user.Nama,
			"no_telp":       user.NoTelp,
			"tanggal_lahir": user.TanggalLahir,
			"jenis_kelamin": user.JenisKelamin,
			"pekerjaan":     user.Pekerjaan,
			"email":         user.Email,
			"id_provinsi":   user.IDProvinsi,
			"id_kota":       user.IDKota,
		},
		"toko": fiber.Map{
			"id":        toko.ID,
			"nama_toko": toko.NamaToko,
			"url_foto":  toko.UrlFoto,
		},
	})
}

func Login(c *fiber.Ctx) error {
	godotenv.Load()

	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		return c.Status(500).JSON(fiber.Map{"error": "JWT secret tidak ditemukan"})
	}

	var loginData struct {
		NoTelp    string `json:"no_telp"`
		KataSandi string `json:"kata_sandi"`
	}
	if err := c.BodyParser(&loginData); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Gagal parsing data"})
	}

	var user models.User
	if err := database.DB.Where("no_telp = ?", loginData.NoTelp).First(&user).Error; err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Nomor telepon tidak terdaftar"})
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.KataSandi), []byte(loginData.KataSandi))
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Kata sandi salah"})
	}

	expirationTime := time.Now().Add(30 * 24 * time.Hour)
	claims := jwt.MapClaims{
		"id":       user.ID,
		"nama":     user.Nama,
		"no_telp":  user.NoTelp,
		"is_admin": user.IsAdmin,
		"exp":      expirationTime.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal membuat token"})
	}

	return c.JSON(fiber.Map{
		"message": "Login berhasil",
		"token":   tokenString,
		"user": fiber.Map{
			"id":      user.ID,
			"nama":    user.Nama,
			"no_telp": user.NoTelp,
			"email":   user.Email,
		},
	})
}

func GetUser(c *fiber.Ctx) error {
	// Ambil userID dari token
	userIDFloat, ok := c.Locals("userID").(float64)
	if !ok {
		return c.Status(400).JSON(fiber.Map{"error": "Gagal mendapatkan userID dari token"})
	}
	userID := uint(userIDFloat)

	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User tidak ditemukan"})
	}

	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	// Ambil userID dari token
	userIDFloat, ok := c.Locals("userID").(float64)
	if !ok {
		return c.Status(400).JSON(fiber.Map{"error": "Gagal mendapatkan userID dari token"})
	}
	userID := uint(userIDFloat)

	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User tidak ditemukan"})
	}

	// Ambil data dari request body
	var updateData struct {
		Nama         string `json:"nama"`
		KataSandi    string `json:"kata_sandi"`
		NoTelp       string `json:"no_telp"`
		TanggalLahir string `json:"tanggal_Lahir"`
		Pekerjaan    string `json:"pekerjaan"`
		Email        string `json:"email"`
		IDProvinsi   string `json:"id_provinsi"`
		IDKota       string `json:"id_kota"`
	}

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Gagal parsing data"})
	}

	// Jika ada perubahan kata sandi, hash ulang
	if updateData.KataSandi != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(updateData.KataSandi), bcrypt.DefaultCost)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Gagal mengenkripsi password"})
		}
		user.KataSandi = string(hashedPassword)
	}

	// Konversi format tanggal lahir
	if updateData.TanggalLahir != "" {
		parsedDate, err := time.Parse("02/01/2006", updateData.TanggalLahir)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Format tanggal harus DD/MM/YYYY"})
		}
		user.TanggalLahir = parsedDate.Format("2006-01-02")
	}

	// Update data user
	user.Nama = updateData.Nama
	user.NoTelp = updateData.NoTelp
	user.Pekerjaan = updateData.Pekerjaan
	user.Email = updateData.Email
	user.IDProvinsi = updateData.IDProvinsi
	user.IDKota = updateData.IDKota

	if err := database.DB.Save(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal menyimpan perubahan"})
	}

	return c.JSON(fiber.Map{"message": "Data user berhasil diperbarui", "user": user})
}
