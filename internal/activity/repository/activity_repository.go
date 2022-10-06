package repository

import (
	"golangtesting/internal/domain"
	"time"

	"gorm.io/gorm"
)

type ActivityRepository struct {
	db *gorm.DB
}

func NewActivityRepository(db *gorm.DB) *ActivityRepository {
	return &ActivityRepository{db}
}

func (ar ActivityRepository) Getdata() ([]domain.Activity, error) {
	var activity []domain.Activity
	err := ar.db.Find(&activity).Error

	return activity, err
}

func (ar ActivityRepository) Getdatabyid(id string) (domain.Activity, error) {
	var activity domain.Activity
	err := ar.db.First(&activity, id).Error

	return activity, err
}

func (ar ActivityRepository) Create(activity domain.Activity) (domain.Activity, error) {
	activity.DeletedAt = time.Time{}
	err := ar.db.Create(&activity).Error
	return activity, err
}

func (ar ActivityRepository) Delete(id string) (domain.Activity, error) {
	var activity domain.Activity
	err := ar.db.Delete(&activity, id).Error
	return activity, err
}

func (ar ActivityRepository) Update(activities domain.Activity, id string) (domain.Activity, error) {
	err := ar.db.Model(&domain.Activity{}).Where("id = ?", id).Updates(domain.Activity{
		Email:     activities.Email,
		Title:     activities.Title,
		CreatedAt: activities.CreatedAt,
		UpdatedAt: activities.UpdatedAt,
	}).Error

	return activities, err

}
