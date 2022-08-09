package controllers

import (
	"log"
	"mitramas_test/auth"
	"mitramas_test/helpers"
	"mitramas_test/models"
	"mitramas_test/services"
	"mitramas_test/validations"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

type ActivityController struct {
	service services.ActivityService
}

func NewActivityController(service services.ActivityService) *ActivityController {
	return &ActivityController{service}
}

func GetUserId(c *gin.Context) (int, error) {
	userId, err := c.Request.Cookie("user_id")
	if err != nil {
		return 0, err
	}

	userIdInt, err := strconv.Atoi(userId.Value)
	if err != nil {
		helpers.NewHandlerResponse("Error convert data", nil).BadRequest(c)
		return 0, err
	}
	return userIdInt, nil
}

func GetCheckInId(c *gin.Context) (int, error) {
	checkInId, err := c.Request.Cookie("check_in_id")
	if err != nil {
		return 0, err
	}

	checkInIdInt, err := strconv.Atoi(checkInId.Value)
	if err != nil {
		helpers.NewHandlerResponse("Error convert data", nil).BadRequest(c)
		return 0, err
	}
	return checkInIdInt, nil
}

func (activityController *ActivityController) CheckInController(c *gin.Context) {
	var checkInId int
	userId, err := auth.ParseJWTClaims(c)
	if err != nil {
		log.Printf("Error read claims with err: %s", err)
		helpers.NewHandlerResponse("Error read claims", nil).BadRequest(c)
		return
	}

	checkInId, err = activityController.service.CheckIn(*userId)
	if err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return
	}

	checkInStr := strconv.Itoa(checkInId)

	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "check_in_id",
		Value:   checkInStr,
		Expires: time.Now().Add(30 * time.Hour),
		Path:    "/",
		// Local
		SameSite: 2,
		HttpOnly: true,
	})

	helpers.NewHandlerResponse("Successfully check-in", nil).SuccessCreate(c)
}

func (activityController *ActivityController) CreateActivityController(c *gin.Context) {
	var activity models.Activity
	checkInId, err := GetCheckInId(c)
	if err != nil {
		helpers.NewHandlerResponse("Please check-in first", nil).BadRequest(c)
		return
	}
	if err := c.ShouldBindJSON(&activity); err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).BadRequest(c)
		return
	}

	if err := validations.DoValidation(&activity); err != nil {
		helpers.NewHandlerValidationResponse(err, nil).BadRequest(c)
		return
	}

	if err := activityController.service.CreateActivity(checkInId, &activity); err != nil {
		sqlErr := err.(*mysql.MySQLError)
		if sqlErr.Number == 1062 {
			helpers.NewHandlerResponse(sqlErr.Message, nil).Failed(c)
			return
		}
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return
	}
	helpers.NewHandlerResponse("Successfully create activity", nil).SuccessCreate(c)
}

func (activityController *ActivityController) ReadActivityController(c *gin.Context) {
	userId, err := auth.ParseJWTClaims(c)
	if err != nil {
		log.Printf("Error read claims with err: %s", err)
		helpers.NewHandlerResponse("Error read claims", nil).BadRequest(c)
		return
	}

	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	activities, err := activityController.service.ReadActivity(*userId, startDate, endDate)
	if err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return
	}
	helpers.NewHandlerResponse("Successfully get activities", activities).Success(c)
}

func (activityController *ActivityController) UpdateActivityController(c *gin.Context) {
	var activity models.Activity

	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		helpers.NewHandlerResponse("Error convert data", nil).BadRequest(c)
		return
	}

	if err := c.ShouldBindJSON(&activity); err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).BadRequest(c)
		return
	}

	if err := validations.DoValidation(&activity); err != nil {
		helpers.NewHandlerValidationResponse(err, nil).BadRequest(c)
		return
	}

	if err := activityController.service.UpdateActivity(id, &activity); err != nil {
		if err.Error() == "id activity not fount" {
			helpers.NewHandlerResponse("Id activity not fount", nil).Failed(c)
			return
		}
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return
	}
	helpers.NewHandlerResponse("Successfully update activity", nil).Success(c)
}

func (activityController *ActivityController) DeleteActivityController(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		helpers.NewHandlerResponse("Error convert data", nil).BadRequest(c)
		return
	}

	if err := activityController.service.DeleteActivity(id); err != nil {
		if err.Error() == "id activity not fount" {
			helpers.NewHandlerResponse("Id activity not fount", nil).Failed(c)
			return
		}
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return
	}

	helpers.NewHandlerResponse("Successfully delete activity", nil).Success(c)
}

func (activityController *ActivityController) ReadAttendanceController(c *gin.Context) {
	userId, err := auth.ParseJWTClaims(c)
	if err != nil {
		log.Printf("Error read claims with err: %s", err)
		helpers.NewHandlerResponse("Error read claims", nil).BadRequest(c)
		return
	}

	attendances, err := activityController.service.ReadAttendance(*userId)
	if err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return
	}
	helpers.NewHandlerResponse("Successfully get attendances", attendances).Success(c)
}

func (activityController *ActivityController) CheckOutController(c *gin.Context) {
	checkInId, err := GetCheckInId(c)
	if err != nil {
		helpers.NewHandlerResponse("Please check-in first", nil).BadRequest(c)
		return
	}

	if err := activityController.service.CheckOut(checkInId); err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "check_in_id",
		Value:   "",
		Expires: time.Unix(0, 0),
		Path:    "/",
		// Local
		SameSite: 2,
		HttpOnly: true,
	})

	helpers.NewHandlerResponse("Successfully checkout", nil).Success(c)
}
