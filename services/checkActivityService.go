package services

import (
	"log"
	"mitramas_test/models"
	"mitramas_test/repositories"
	"time"
)

type ActivityService interface {
	CheckIn(userId int) (int, error)
	CreateActivity(checkInId int, activity *models.Activity) error
	ReadActivity(userId int, startDate, endDate string) (*[]models.Activity, error)
	UpdateActivity(id int, activity *models.Activity) error
	DeleteActivity(Id int) error
	ReadAttendance(userId int) (*[]models.Attendance, error)
}

type activityService struct {
	activityRepository repositories.ActivityRepository
}

func NewActivityService(repository repositories.ActivityRepository) *activityService {
	return &activityService{repository}
}

func (s *activityService) CheckIn(userId int) (int, error) {
	var (
		checkInId int
		err       error
	)
	Checkin := models.CheckIn{
		UserId:      userId,
		DateCheckIn: time.Now(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	checkInId, err = s.activityRepository.CheckIn(&Checkin)
	if err != nil {
		log.Printf("Error create checkin to database with err: %s", err)
		return checkInId, err
	}

	return checkInId, nil
}

func (s *activityService) CreateActivity(checkInId int, activity *models.Activity) error {
	activityCreate := models.Activity{
		CheckInId:   checkInId,
		Description: activity.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	if err := s.activityRepository.CreateActivity(&activityCreate); err != nil {
		log.Printf("Error create activity to database with err: %s", err)
		return err
	}
	return nil
}

func (s *activityService) ReadActivity(userId int, startDate, endDate string) (*[]models.Activity, error) {
	activities, err := s.activityRepository.ReadActivity(userId, startDate, endDate)
	if err != nil {
		log.Printf("Error get data activities with err: %s", err)
		return nil, err
	}

	return activities, nil
}

func (s *activityService) UpdateActivity(id int, activity *models.Activity) error {
	activityUpdate := models.Activity{
		Id:          id,
		Description: activity.Description,
		UpdatedAt:   time.Now(),
	}

	if err := s.activityRepository.UpdateActivity(&activityUpdate); err != nil {
		log.Printf("Error update activity to database with err: %s", err)
		return err
	}

	return nil
}

func (s *activityService) DeleteActivity(id int) error {
	if err := s.activityRepository.DeleteActivity(id); err != nil {
		log.Printf("Error delete data activity with err: %s", err)
		return err
	}

	return nil
}

func (s *activityService) ReadAttendance(userId int) (*[]models.Attendance, error) {
	attendances, err := s.activityRepository.ReadAttendance(userId)
	if err != nil {
		log.Printf("Error get data attendances with err: %s", err)
		return nil, err
	}

	return attendances, nil
}
