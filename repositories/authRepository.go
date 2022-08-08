package repositories

import (
	"database/sql"
	"log"
	"mitramas_test/models"
)

type AuthRepository interface {
	Register(user *models.UserRegister) error
	Login(user *models.UserLogin) (*models.UserLogin, error)
}

type authRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *authRepository {
	return &authRepository{db}
}

func (r *authRepository) Register(user *models.UserRegister) error {
	db := r.db
	sqlStatement := `INSERT INTO users (user_name, password, email, created_at, updated_at) VALUES (?, ?, ?, ?, ?)`
	_, err := db.Exec(sqlStatement, &user.UserName, &user.Password, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		log.Printf("Error register user to database with err: %s", err)
		return err
	}
	return nil
}

func (r *authRepository) Login(user *models.UserLogin) (*models.UserLogin, error) {
	db := r.db
	userData := models.UserLogin{}

	sqlInfo := `SELECT id, email, password FROM users WHERE email = ?`
	err := db.QueryRow(sqlInfo, user.Email).Scan(&userData.Id, &userData.Email, &userData.Password)
	if err != nil {
		log.Printf("User not found with err: %s", err)
		return nil, err
	}

	return &userData, err
}
