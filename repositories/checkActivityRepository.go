package repositories

import (
	"database/sql"
	"fmt"
	"log"
	"mitramas_test/models"
	"time"
)

type ActivityRepository interface {
	CheckIn(checkIn *models.CheckIn) (int, error)
	CreateActivity(activity *models.Activity) error
	ReadActivity(userId int, startDate, endDate string) (*[]models.Activity, error)
	UpdateActivity(activity *models.Activity) error
	DeleteActivity(id int) error
	ReadAttendance(userId int) (*[]models.Attendance, error)
	CheckOut(checkOut *models.CheckOut) error
}

type activityRepository struct {
	db *sql.DB
}

func NewActivityRepository(db *sql.DB) *activityRepository {
	return &activityRepository{db}
}

func (r *activityRepository) CheckIn(checkIn *models.CheckIn) (int, error) {
	db := r.db
	var checkInId int
	sqlStatement := `INSERT INTO check_ins (date_check_in, user_id, created_at, updated_at) VALUES (?, ?, ?, ?)`

	_, err := db.Exec(sqlStatement, &checkIn.DateCheckIn, &checkIn.UserId, &checkIn.CreatedAt, &checkIn.UpdatedAt)

	if err != nil {
		log.Printf("Error create check in to database with err: %s", err)
		return checkInId, err
	}

	// GET LAST START SHIFT
	getLastIdCheckIn := "SELECT LAST_INSERT_ID() FROM check_ins"
	err = db.QueryRow(getLastIdCheckIn).Scan(&checkInId)
	if err != nil {
		log.Printf("Error get data last check in id with err: %s", err)
		return checkInId, err
	}

	return checkInId, nil
}

func (r *activityRepository) CreateActivity(activity *models.Activity) error {
	db := r.db
	sqlStatement := `INSERT INTO activities (check_in_id, description, created_at, updated_at) VALUES (?, ?, ?, ?)`

	_, err := db.Exec(sqlStatement, &activity.CheckInId, &activity.Description, &activity.CreatedAt, &activity.UpdatedAt)

	if err != nil {
		log.Printf("Error create activity to database with err: %s", err)
		return err
	}

	return nil
}

func (r *activityRepository) ReadActivity(userId int, startDate, endDate string) (*[]models.Activity, error) {
	db := r.db
	var result = []models.Activity{}
	t := time.Now()
	formatted := fmt.Sprintf("%d-%02d-%02d", t.Year(), t.Month(), t.Day())
	var queryDate string

	if startDate == "" || endDate == "" {
		queryDate = fmt.Sprintf(` DATE_FORMAT(activities.created_at, '%%Y-%%m-%%d') = '%s'`, formatted)
	} else {
		startDate += " 0:0:0"
		endDate += " 23:59:59"
		queryDate = fmt.Sprintf(` activities.created_at BETWEEN '%s' AND '%s'`, startDate, endDate)
	}

	sqlStatement := fmt.Sprintf(`SELECT activities.id, activities.check_in_id, activities.description, activities.created_at,
	activities.updated_at, DATE_FORMAT(activities.created_at, '%%d %%M %%Y, %%H:%%i') as date FROM activities
	LEFT JOIN check_ins ON activities.check_in_id=check_ins.id
	WHERE check_ins.user_id = ? AND %s
	ORDER BY activities.created_at DESC`, queryDate)
	rows, err := db.Query(sqlStatement, userId)
	if err != nil {
		log.Printf("Error get data activities with err: %s", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var activity = models.Activity{}

		err := rows.Scan(&activity.Id, &activity.CheckInId, &activity.Description, &activity.CreatedAt, &activity.UpdatedAt, &activity.Date)
		if err != nil {
			log.Printf("Error get data customer reedem with err: %s", err)
			return nil, err
		}

		result = append(result, activity)
	}

	return &result, nil
}

func (r *activityRepository) UpdateActivity(activity *models.Activity) error {
	db := r.db
	sqlUpdate := `UPDATE activities SET description = ?, updated_at = ? WHERE id = ?`
	res, err := db.Exec(sqlUpdate, &activity.Description, &activity.UpdatedAt, &activity.Id)
	if err != nil {
		log.Printf("Error update activity to database with err: %s", err)
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error update activity to database with err: %s", err)
		return err
	}
	if count == 0 {
		log.Printf("Id activity not fount")
		return fmt.Errorf("id activity not fount")
	}

	return nil
}

func (r *activityRepository) DeleteActivity(id int) error {
	db := r.db
	sqlDelete := `DELETE FROM activities WHERE id = ?`
	res, err := db.Exec(sqlDelete, id)
	if err != nil {
		log.Printf("Error delete data activity to database with err: %s", err)
		return err
	}

	count, errs := res.RowsAffected()
	if errs != nil {
		log.Printf("Error delete data activity to database with err: %s", err)
		return errs
	}
	if count == 0 {
		log.Printf("Id activity not fount")
		return fmt.Errorf("id activity not fount")
	}

	return nil
}

func (r *activityRepository) ReadAttendance(userId int) (*[]models.Attendance, error) {
	db := r.db
	var result = []models.Attendance{}
	sqlStatement := `SELECT check_ins.id, check_ins.user_id, DATE_FORMAT(check_ins.date_check_in, '%d %M %Y, %H:%i') as date_check_in,
	users.user_name as user_name, DATE_FORMAT(check_outs.date_check_out, '%d %M %Y, %H:%i') as date_check_out
	FROM check_ins
	LEFT JOIN users ON check_ins.user_id=users.id
	LEFT JOIN check_outs ON check_outs.check_in_id=check_ins.id
	WHERE check_ins.user_id = ?
	ORDER BY check_ins.date_check_in DESC`
	rows, err := db.Query(sqlStatement, userId)
	if err != nil {
		log.Printf("Error get data attendances with err: %s", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var attendance = models.Attendance{}

		err := rows.Scan(&attendance.Id, &attendance.UserId, &attendance.DateCheckIn, &attendance.UserName, &attendance.DateCheckOut)
		if err != nil {
			log.Printf("Error get data attendances with err: %s", err)
			return nil, err
		}
		result = append(result, attendance)
	}

	return &result, nil
}

func (r *activityRepository) CheckOut(checkOut *models.CheckOut) error {
	db := r.db
	sqlStatement := `INSERT INTO check_outs (check_in_id, date_check_out, created_at, updated_at) VALUES (?, ?, ?, ?)`

	_, err := db.Exec(sqlStatement, &checkOut.CheckInId, &checkOut.DateCheckOut, &checkOut.CreatedAt, &checkOut.UpdatedAt)

	if err != nil {
		log.Printf("Error create check out to database with err: %s", err)
		return err
	}

	return nil
}
