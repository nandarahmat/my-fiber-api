package models

import "time"

type Toko struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	IDUser    uint      `json:"id_user"`
	NamaToko  string    `json:"nama_toko"`
	UrlFoto   string    `json:"url_foto"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (Toko) TableName() string {
	return "toko"
}
