package service

import (
	"errors"
	"time"

	"Mini_Project-name/internal/app/repository"
	"Mini_Project/internal/app/model"
)

type MeteranService struct {
	meteranRepo repository.MeteranRepository
	userRepo    repository.UserRepository
}

func NewMeteranService(meteranRepo repository.MeteranRepository, userRepo repository.UserRepository) *MeteranService {
	return &MeteranService{
		meteranRepo: meteranRepo,
		userRepo:    userRepo,
	}
}

func (s *MeteranService) CreateMeteran(req *model.MeteranRequest, userID uuid.UUID) (*model.Meteran, error) {
	user, err := s.userRepo.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	var prevMeteran *model.Meteran
	if user.LastMeteranID != nil {
		prevMeteran, err = s.meteranRepo.GetMeteranByID(*user.LastMeteranID)
		if err != nil {
			return nil, err
		}
	}

	meteran := &model.Meteran{
		UserID:    userID,
		Meteran:   req.Meteran,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if prevMeteran != nil {
		if req.Meteran < prevMeteran.Meteran {
			return nil, errors.New("current meteran should be greater than previous meteran")
		}

		usage := req.Meteran - prevMeteran.Meteran
		meteran.Usage = &usage
		meteran.TotalPrice = float64(usage) * 5000
	}

	err = s.meteranRepo.CreateMeteran(meteran)
	if err != nil {
		return nil, err
	}

	user.LastMeteranID = &meteran.ID
	err = s.userRepo.UpdateUser(user)
	if err != nil {
		return nil, err
	}

	return meteran, nil
}

func (s *MeteranService) GetMeteranList(userID uuid.UUID) ([]model.Meteran, error) {
	return s.meteranRepo.GetMeteranList(userID)
}
