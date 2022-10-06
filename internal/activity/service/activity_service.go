package service

import (
	"errors"
	"fmt"
	"golangtesting/internal/domain"
	"strconv"
	"time"
)

type ActivityRepo interface {
	Getdata() ([]domain.Activity, error)
	Getdatabyid(id string) (domain.Activity, error)
	Create(activity domain.Activity) (domain.Activity, error)
	Delete(id string) (domain.Activity, error)
	Update(activities domain.Activity, id string) (domain.Activity, error)
}

type ActivityService struct {
	ActivitySer ActivityRepo
}

func NewActivityService(ar ActivityRepo) *ActivityService {
	return &ActivityService{ar}
}

func (as ActivityService) GetdataService() ([]domain.Activity, int, error) {
	activity, err := as.ActivitySer.Getdata()
	if err != nil {
		return nil, 500, errors.New("Error get data")
	}
	return activity, 200, nil
}

func (as ActivityService) GetdataByidService(id string) (domain.Activity, int, error) {
	activity, err := as.ActivitySer.Getdatabyid(id)
	if err != nil {
		msg := fmt.Sprintf("Activity with ID %s Not Found", id)
		return activity, 404, errors.New(msg)
	}
	return activity, 200, nil

}

func (as ActivityService) CreateDataService(activites domain.Activity) (domain.Activity, int, error) {
	activity, err := as.ActivitySer.Create(activites)

	if err != nil {
		return activity, 500, errors.New("Create error from server")
	}
	return activity, 201, nil

}

func (as ActivityService) DeleteDataService(id string) (domain.Activity, int, error) {

	activity, err := as.ActivitySer.Getdatabyid(id)
	if err != nil {
		msg := fmt.Sprintf("Activity with ID %s Not Found", id)
		return activity, 404, errors.New(msg)
	} else {
		activity, _ := as.ActivitySer.Delete(id)

		return activity, 200, errors.New("Success")
	}

}

func (as ActivityService) UpdateService(activities domain.Activity, id string) (domain.Activity, int, error) {
	activity, err := as.ActivitySer.Getdatabyid(id)
	if err != nil {
		msg := fmt.Sprintf("Activity with ID %s Not Found", id)
		return activity, 404, errors.New(msg)
	} else {
		activity, _ := as.ActivitySer.Update(activities, id)
		id_primary, _ := strconv.Atoi(id)
		activity.ID = id_primary
		activity.CreatedAt = time.Now()
		activity.UpdatedAt = time.Now()
		return activity, 200, nil
	}
}
