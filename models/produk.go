package models

import (
	"time"
)

// Model Produk
type Product struct {
	ID            uint           `json:"id" gorm:"primaryKey"`
	NamaProduk    string         `json:"nama_produk"`
	Slug          string         `json:"slug"`
	HargaReseller string         `json:"harga_reseller"`
	HargaKonsumen string         `json:"harga_konsumen"`
	Stok          int            `json:"stok"`
	Deskripsi     string         `json:"deskripsi"`
	IDToko        uint           `json:"id_toko"`
	IDCategory    uint           `json:"id_category"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	Images        []ProductImage `json:"images" gorm:"foreignKey:IDProduct"`
}

// Model Foto Produk
type ProductImage struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	IDProduct uint      `json:"id_produk" gorm:"column:id_produk"`
	URL       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Product) TableName() string {
	return "produk"
}

func (ProductImage) TableName() string {
	return "foto_produk"
}
