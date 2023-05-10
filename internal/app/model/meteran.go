package model

import (
	"time"
)

type Meteran struct {
	ID         uint      `gorm:"primary_key"`
	UserID     string    `gorm:"size:36"`
	Bulan      int       `gorm:"not null"`
	Tahun      int       `gorm:"not null"`
	MeterAwal  int       `gorm:"not null"`
	MeterAkhir int       `gorm:"not null"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

type MeteranRequest struct {
	Bulan      int `json:"bulan"`
	Tahun      int `json:"tahun"`
	MeterAwal  int `json:"meter_awal"`
	MeterAkhir int `json:"meter_akhir"`
}
