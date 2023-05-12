package models

import (
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
    Nama    string `json:"nama"`
    Email   string `json:"email"`
    Umur    int    `json:"umur"`
	Laporan []Bill `gorm:"foreignKey:PelangganID"`
}

type Bill struct {
	gorm.Model
    TotalDebit  int    `json:"total_debit"`
    HargaAir    int    `json:"harga_air"`
    PelangganID uint   `json:"pelanggan_id"`
    Pelanggan   Customer `gorm:"foreignKey:PelangganID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

}