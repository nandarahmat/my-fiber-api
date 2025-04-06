package models

import "time"

type Alamat struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	IDUser       uint      `json:"id_user"`
	Judul        string    `json:"judul"`
	NamaPenerima string    `json:"nama_penerima"`
	NoTelp       string    `json:"no_telp"`
	DetailAlamat string    `json:"detail_alamat"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (Alamat) TableName() string {
	return "alamat"
}
