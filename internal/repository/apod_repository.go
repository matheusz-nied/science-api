package repository

import (
	db "nied-science/internal/database"
	model "nied-science/internal/model"
)

type APODRepository interface {
	SaveAPOD(apod *model.APOD) error
	GetAPODByDate(date string) (*model.APOD, error)
	SaveAPODs(apods []model.APOD) error
	GetAPODsPaginated(page int, pageSize int) ([]model.APOD, int64, error)
}

type apodRepository struct{}

func NewAPODRepository() APODRepository {
	return &apodRepository{}
}

func (r *apodRepository) SaveAPOD(apod *model.APOD) error {
	return db.DB.Create(apod).Error
}

func (r *apodRepository) GetAPODByDate(date string) (*model.APOD, error) {
	var apod model.APOD
	err := db.DB.Where("date = ?", date).First(&apod).Error
	return &apod, err
}

func (r *apodRepository) SaveAPODs(apods []model.APOD) error {
	return db.DB.Create(&apods).Error
}
func (r *apodRepository) GetAPODsPaginated(page int, pageSize int) ([]model.APOD, int64, error) {
	var apods []model.APOD
	var totalItems int64

	offset := (page - 1) * pageSize
	err := db.DB.Order("date DESC").Limit(pageSize).Offset(offset).Find(&apods).Count(&totalItems).Error
	if err != nil {
		return nil, 0, err
	}

	db.DB.Model(&model.APOD{}).Count(&totalItems)

	return apods, totalItems, nil
}
