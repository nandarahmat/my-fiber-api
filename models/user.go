package models

import (
	"time"
)

type User struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Nama         string    `json:"nama"`
	KataSandi    string    `json:"kata_sandi"`
	NoTelp       string    `gorm:"unique" json:"no_telp"`
	TanggalLahir string    `json:"tanggal_lahir"`
	JenisKelamin string    `json:"jenis_kelamin"`
	Pekerjaan    string    `json:"pekerjaan"`
	Tentang      string    `json:"tentang"`
	Email        string    `gorm:"unique" json:"email"`
	IDProvinsi   string    `json:"id_provinsi"`
	IDKota       string    `json:"id_kota"`
	IsAdmin      int       `json:"is_admin"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (User) TableName() string {
	return "user"
}
