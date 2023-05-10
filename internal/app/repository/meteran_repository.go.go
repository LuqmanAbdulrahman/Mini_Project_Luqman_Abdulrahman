package repository

import (
	"errors"

	"gorm.io/gorm"

	"Mini_Project/internal/app/model"
)

type MeteranRepository struct {
	db *gorm.DB
}

func NewMeteranRepository(db *gorm.DB) *MeteranRepository {
	return &MeteranRepository{db: db}
}

func (r *MeteranRepository) CreateMeteran(meteran *model.Meteran) error {
	if err := r.db.Create(meteran).Error; err != nil {
		return err
	}
	return nil
}

func (r *MeteranRepository) FindMeteranByUserIDAndBulanAndTahun(userID string, bulan, tahun int) (*model.Meteran, error) {
	var meteran model.Meteran
	if err := r.db.Where("user_id = ? AND bulan = ? AND tahun = ?", userID, bulan, tahun).First(&meteran).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &meteran, nil
}

func (r *MeteranRepository) UpdateMeteran(meteran *model.Meteran) error {
	if err := r.db.Save(meteran).Error; err != nil {
		return err
	}
	return nil
}

func (r *MeteranRepository) GetTotalPemakaian(userID string, bulan, tahun int) (int, error) {
	var totalPemakaian int
	if err := r.db.Raw("SELECT SUM(meter_akhir - meter_awal) AS total_pemakaian FROM meteran WHERE user_id = ? AND bulan = ? AND tahun = ?", userID, bulan, tahun).Scan(&totalPemakaian).Error; err != nil {
		return 0, err
	}
	return totalPemakaian, nil
}

func (r *MeteranRepository) GetTotalTagihan(userID string, bulan, tahun int, hargaPerKubik int) (int, error) {
	totalPemakaian, err := r.GetTotalPemakaian(userID, bulan, tahun)
	if err != nil {
		return 0, err
	}
	return totalPemakaian * hargaPerKubik, nil
}
