package controllers

import (
	"database/sql"
	"errors"
	"fmt"
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
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {
	service services.AuthService
}

func NewAuthController(service services.AuthService) *AuthController {
	return &AuthController{service}
}

func (authController *AuthController) RegisterController(c *gin.Context) {
	var user models.UserRegister

	if err := c.ShouldBindJSON(&user); err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).BadRequest(c)
		return
	}

	if err := validations.DoValidation(&user); err != nil {
		helpers.NewHandlerValidationResponse(err, nil).BadRequest(c)
		return
	}

	err := authController.service.Register(&user)
	if err != nil {
		sqlErr := err.(*mysql.MySQLError)
		if sqlErr.Number == 1062 {
			helpers.NewHandlerResponse(sqlErr.Message, nil).Failed(c)
			return
		}
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return
	}
	helpers.NewHandlerResponse("Successfully register", nil).SuccessCreate(c)
}

func (authController *AuthController) LoginController(c *gin.Context) {
	var user models.UserLogin
	if err := c.ShouldBindJSON(&user); err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).BadRequest(c)
		return
	}

	if err := validations.DoValidation(&user); err != nil {
		helpers.NewHandlerValidationResponse(err, nil).BadRequest(c)
		return
	}

	userData, err := authController.service.Login(&user)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			helpers.NewHandlerResponse("Email Not Found", nil).Failed(c)
			return
		}
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return
	}

	errHash := bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(user.Password))
	if errHash != nil {
		fmt.Printf("Password Incorrect with err: %s\n", errHash)
		helpers.NewHandlerResponse("Password Incorrect", nil).BadRequest(c)
		return
	}

	tokenString, err := auth.GenerateJWT(userData)
	if err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		// c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	id := strconv.Itoa(userData.Id)

	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: time.Now().Add(30 * time.Hour),
		Path:    "/",
		// Local
		SameSite: 2,
		HttpOnly: true,
	})
	// userIdCookie := &http.Cookie{
	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "user_id",
		Value:   id,
		Expires: time.Now().Add(30 * time.Hour),
		Path:    "/",
		// Local
		SameSite: 2,
		HttpOnly: true,
	})

	helpers.NewHandlerResponse("Successfully Login", nil).Success(c)
}

func (authController *AuthController) LogoutController(c *gin.Context) {
	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "token",
		Value:   "",
		Expires: time.Unix(0, 0),
		Path:    "/",
		// Local
		SameSite: 2,
		HttpOnly: true,
	})
	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "user_id",
		Value:   "",
		Expires: time.Unix(0, 0),
		Path:    "/",
		// Local
		SameSite: 2,
		HttpOnly: true,
	})
	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "check_in_id",
		Value:   "",
		Expires: time.Unix(0, 0),
		Path:    "/",
		// Local
		SameSite: 2,
		HttpOnly: true,
	})
	helpers.NewHandlerResponse("Successfully logout", nil).Success(c)
}
