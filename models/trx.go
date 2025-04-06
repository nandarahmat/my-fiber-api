package models

import (
	"time"
)

type Trx struct {
	ID               uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID           uint      `json:"id_user" gorm:"column:id_user;not null;index"`
	AlamatPengiriman uint      `json:"alamat_pengiriman" gorm:"not null;index"`
	HargaTotal       int       `json:"harga_total" gorm:"not null"`
	KodeInvoice      string    `json:"kode_invoice" gorm:"size:255;not null;unique"`
	MethodBayar      string    `json:"method_bayar" gorm:"size:255;not null"`
	CreatedAt        time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt        time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (Trx) TableName() string {
	return "trx"
}

// LogProduk model
type LogProduk struct {
	ID            uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	IDProduk      uint      `json:"id_produk" gorm:"not null;index"`
	NamaProduk    string    `json:"nama_produk" gorm:"not null;size:255"`
	Slug          string    `json:"slug" gorm:"not null;size:255;unique"`
	HargaReseller int       `json:"harga_reseller" gorm:"not null"`
	HargaKonsumen int       `json:"harga_konsumen" gorm:"not null"`
	Deskripsi     string    `json:"deskripsi" gorm:"type:text"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	IDToko        uint      `json:"id_toko" gorm:"not null;index"`
	IDCategory    uint      `json:"id_category" gorm:"not null;index"`
}

func (LogProduk) TableName() string {
	return "log_produk"
}

// DetailTrx model
type DetailTrx struct {
	ID          uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	IDTrx       uint      `json:"id_trx" gorm:"not null;index"`
	IDLogProduk uint      `json:"id_log_produk" gorm:"not null;index"`
	IDToko      uint      `json:"id_toko" gorm:"not null;index"`
	Kuantitas   int       `json:"kuantitas" gorm:"not null"`
	HargaTotal  int       `json:"harga_total" gorm:"not null"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (DetailTrx) TableName() string {
	return "detail_trx"
}
