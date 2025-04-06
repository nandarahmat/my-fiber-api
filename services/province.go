package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Province struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Regency struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func GetProvinces(c *fiber.Ctx) error {
	url := "https://www.emsifa.com/api-wilayah-indonesia/api/provinces.json"

	resp, err := http.Get(url)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal mengambil data provinsi"})
	}
	defer resp.Body.Close()

	var provinces []Province
	if err := json.NewDecoder(resp.Body).Decode(&provinces); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal decode data"})
	}

	return c.JSON(provinces)
}

func GetDetailProvinces(c *fiber.Ctx) error {
	provinceID := c.Params("id")

	url := fmt.Sprintf("https://www.emsifa.com/api-wilayah-indonesia/api/regencies/%s.json", provinceID)

	resp, err := http.Get(url)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal mengambil data detail provinsi"})
	}
	defer resp.Body.Close()

	var provinces []Province
	if err := json.NewDecoder(resp.Body).Decode(&provinces); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal decode data"})
	}

	return c.JSON(provinces)
}

func GetCity(c *fiber.Ctx) error {
	cityID := c.Params("id")

	url := fmt.Sprintf("https://www.emsifa.com/api-wilayah-indonesia/api/districts/%s.json", cityID)

	resp, err := http.Get(url)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal mengambil data kecamatan"})
	}
	defer resp.Body.Close()

	var regencies []Regency
	if err := json.NewDecoder(resp.Body).Decode(&regencies); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Data kecamatan tidak tersedia"})
	}

	return c.JSON(regencies)
}

func GetDetailCity(c *fiber.Ctx) error {
	cityID := c.Params("id")

	url := fmt.Sprintf("https://www.emsifa.com/api-wilayah-indonesia/api/villages/%s.json", cityID)

	resp, err := http.Get(url)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal mengambil data kelurahan"})
	}
	defer resp.Body.Close()

	var regencies []Regency
	if err := json.NewDecoder(resp.Body).Decode(&regencies); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Data kelurahan tidak tersedia"})
	}

	return c.JSON(regencies)
}
